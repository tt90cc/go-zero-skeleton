FROM golang:1.18-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata && \
apk add openssh-server && apk add openssh-client && \
apk add git && git config --global url."git@gitee.com:".insteadOf "https://gitee.com/"

WORKDIR /app

COPY . .

RUN mkdir -p ~/.ssh && cp ./doc/id_rsa ~/.ssh && chmod 0600 ~/.ssh/id_rsa && ssh-keyscan gitee.com >> ~/.ssh/known_hosts && \
go env -w GOPRIVATE=gitee.com && \
go get github.com/tt90cc/utils \
go mod tidy && go build -ldflags="-s -w" -o main ./main.go

RUN mkdir -p publish/etc && \
cp ./main publish && \
cp -r ./etc/pro/* publish/etc

FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/publish .
RUN chmod +x main

CMD ["./main", "-f", "etc/main.yaml"]