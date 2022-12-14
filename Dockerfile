#build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o main.app 

#final stage
FROM alpine:latest
COPY --from=builder app/main.app /app
ENTRYPOINT /app
LABEL Name=clean_architecture_go Version=1.0
EXPOSE 8080
