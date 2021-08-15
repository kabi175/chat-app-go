FROM golang:alpine as builder

WORKDIR /go/src/app

# Get Reflex for live reload in dev
ENV GO111MODULE=on
RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY .env .

COPY . .

RUN go build -o ./run .

FROM alpine:3.14
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD ["./run"]https://meet.google.com/sfy-jfpf-hvi

