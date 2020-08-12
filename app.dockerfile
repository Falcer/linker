FROM golang:alpine AS builder

RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
COPY vendor vendor
COPY . .
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app .


FROM alpine

WORKDIR /usr/bin
COPY --from=builder /go/bin .
EXPOSE 8080
CMD ["app"]