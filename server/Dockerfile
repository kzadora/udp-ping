#build stage
FROM golang:1-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build server.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

EXPOSE 17335

COPY --from=builder /app/server /app/server

ENTRYPOINT ["/app/server", "17335"]


