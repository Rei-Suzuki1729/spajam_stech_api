FROM golang:latest

WORKDIR /app
COPY ./ /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
EXPOSE 1323


CMD ["go", "run", "server.go"]
