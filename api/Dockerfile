FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app

COPY . .

RUN go build -ldflags="-s -w" -o main ./api/ucenter.go

ARG CONF_ENV

RUN mkdir -p publish/etc && \
    cp ./main publish && \
    cp -r ./api/etc/${CONF_ENV}/* publish/etc


FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/publish .
RUN chmod +x main

CMD ["./main", "-f", "etc/ucenter.yaml"]
