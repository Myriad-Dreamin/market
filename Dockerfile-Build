FROM golang:alpine AS build-env
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

ADD . /market
WORKDIR /market
ENV GO111MODULE=on GOPROXY=https://goproxy.io
RUN apk update && apk add gcc && apk add musl-dev
RUN GOOS=linux GOARCH=amd64 go build -v -o ../server .

FROM alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

WORKDIR /
RUN apk add -U tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY --from=build-env /server /usr/local/bin/server
COPY docker/run.sh /usr/local/bin/run.sh
RUN chmod +x /usr/local/bin/run.sh
EXPOSE 23336
CMD run.sh

