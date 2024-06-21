FROM golang:1.22.4-bullseye as builder

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build ./main/start.go .


FROM alpine:3.20

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /usr/app/main .

EXPOSE 8000

CMD [ "./start" ]
