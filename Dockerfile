# Stage 1: Build the Go application

FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum  ./
RUN go mod download

COPY . .
RUN go build -o indication main.go


# Stage 2: Run the binary in a small container
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/indication .
RUN mkdir -p configs
COPY ./configs/main.yml ./configs/ 

COPY ./configs/log_config.json ./configs/
RUN chmod +x ./indication
EXPOSE 8000

CMD ["./indication"]