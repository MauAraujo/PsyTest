FROM node:lts-alpine AS node-build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY ./ .
RUN npm run build

FROM golang:alpine AS go-build
RUN mkdir /go/src/app
COPY . /go/src/app
WORKDIR /go/src/app
RUN go build -o server .

RUN apk add nginx
RUN mkdir /app
RUN mkdir -p /run/nginx
COPY --from=node-build /app/dist /app
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 3000
EXPOSE 8080

RUN apk add supervisor
RUN mkdir -p /var/log/supervisor
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
CMD ["/usr/bin/supervisord"]