#build stage
FROM golang:1.20-alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
COPY .env .env
ENV DB_USERNAME="${DB_USERNAME}"
ENV DB_PASSWORD="${DB_PASSWORD}"
ENV DB_PORT="${DB_PORT}"
ENV DB_HOST="${DB_HOST}"
ENV DB_NAME="${DB_NAME}"
ENV JWT_SECRET="${JWT_SECRET}"
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /app/.env /app/.env
ENV DB_USERNAME="${DB_USERNAME}"
ENV DB_PASSWORD="${DB_PASSWORD}"
ENV DB_PORT="${DB_PORT}"
ENV DB_HOST="${DB_HOST}"
ENV DB_NAME="${DB_NAME}"
ENV JWT_SECRET="${JWT_SECRET}"
ENTRYPOINT /app
LABEL Name=gosimplecleanrestapi Version=0.0.1
EXPOSE 8080