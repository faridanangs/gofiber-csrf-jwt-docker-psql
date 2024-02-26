FROM golang:1.21.5 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go get -d -v ./...

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

COPY --from=builder /app/app /app/app

CMD ["/app/app"]
