FROM golang:1.15.2 as build_sms-api

ENV CGO_ENABLED 0

RUN mkdir -p /service

COPY go.* /service

WORKDIR /service

RUN go mod download

COPY . .

RUN go build -o sms-webhook-api app/sms-api/*.go

FROM alpine:3.14

RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
    echo "Asia/Seoul" > /etc/timezone

COPY --from=build_sms-api /service/app/sms-api/sms-webhook-api /service/sms-webhook-api
WORKDIR /service
CMD ["./sms-webhook-api"]

#RUN go build -ldflags "-X main.build=${VCS_REF}"

