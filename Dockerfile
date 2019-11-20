FROM node:lts-alpine AS node-build
WORKDIR /go/src/app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM golang:alpine AS go-build
WORKDIR /go/src/app
RUN apk add nginx
RUN mkdir -p /run/nginx
COPY default.conf /etc/nginx/conf.d/default.conf
COPY . .
EXPOSE 8080
EXPOSE 3000
RUN go build -o server .
COPY --from=node-build /go/src/app/dist /usr/share/nginx/html
RUN apk add supervisor
RUN mkdir -p /var/log/supervisor
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
CMD ["/usr/bin/supervisord"]