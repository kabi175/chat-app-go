FROM golang:alpine

WORKDIR /src/messager

# Get Reflex for live reload in dev
ENV GO111MODULE=on
RUN go get github.com/cespare/reflex

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY .env .

COPY . .

RUN go build -o ./app

EXPOSE 8080

CMD ["./app"]