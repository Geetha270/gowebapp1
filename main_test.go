# Build stage
FROM golang:1.22.5 AS base

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main .

# Final stage with Alpine (no need to install bash now)
FROM alpine:latest

# Optional: install bash if you want (not necessary)
# RUN apk add --no-cache bash

WORKDIR /app

COPY --from=base /app/main .
COPY --from=base /app/static ./static

EXPOSE 8080

CMD ["./main"]
