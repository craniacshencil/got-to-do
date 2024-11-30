FROM golang:1.22-alpine AS builder
WORKDIR /app
# No need to download, as i have vendor directory
COPY . .
RUN go build -o bin/got_to_do cmd/main.go


FROM alpine AS production
WORKDIR /build
COPY --from=builder /app/bin/got_to_do .
COPY --from=builder /app/.env.docker .env.local
COPY --from=builder /app/makefile .
COPY --from=builder /app/scripts ./scripts
RUN export APP_ENV="docker"
RUN apk add --no-cache bash make curl
EXPOSE 8080
CMD ["./got_to_do"]
