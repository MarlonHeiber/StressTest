FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o stresstest

FROM alpine
WORKDIR /app
COPY --from=builder /app/stresstest .
ENTRYPOINT ["./stresstest"]
