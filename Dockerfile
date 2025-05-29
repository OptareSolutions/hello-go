
FROM golang:1.20-alpine AS builder
WORKDIR /src
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /hello-go


FROM alpine:latest
RUN addgroup -S app && adduser -S app -G app
USER app
COPY --from=builder /hello-go /hello-go
EXPOSE 8080
ENTRYPOINT ["/hello-go"]