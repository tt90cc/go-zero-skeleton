FROM golang:1.16-alpine3.15 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
ARG CONF_ENV
COPY app/ucenter/api/etc/${CONF_ENV}/* /app/etc
RUN go build -ldflags="-s -w" -o /app/ucenter app/ucenter/api/ucenter.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/ucenter /app/ucenter
COPY --from=builder /app/etc /app/etc

CMD ["./ucenter", "-f", "etc/ucenter.yaml"]
