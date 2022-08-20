FROM golang:latest

WORKDIR /app
COPY ./ /app

ARG DB_HOST
ARG DB_USER
ARG DB_PASSWORD
ARG DB_PORT
ARG DB_NAME

# ENV DB_HOST ${DB_HOST}
# ENV DB_USER ${DB_USER}
# ENV DB_PASSWORD ${DB_PASSWORD}}
# ENV DB_PORT ${DB_PORT}
# ENV DB_NAME ${DB_NAME}

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
EXPOSE 1323


CMD ["go", "run", "server.go"]
