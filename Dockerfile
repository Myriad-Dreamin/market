FROM golang:alpine AS build-env
ADD . /market
WORKDIR /market
ENV GO111MODULE=on GOPROXY=https://goproxy.io
RUN GOOS=linux go build -v -o ../server .

FROM amd64/alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
WORKDIR /
RUN apk add -U tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY --from=build-env /server /usr/local/bin/server
COPY docker/run.sh /usr/local/bin/run.sh
RUN chmod +x /usr/local/bin/run.sh
EXPOSE 23336
CMD [ "run.sh" ]

