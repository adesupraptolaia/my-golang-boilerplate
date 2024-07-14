FROM golang:1.22.1-alpine3.19 as builder

WORKDIR /app

# Copy the source code
COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./main.go

# ############################
# # Stage 2
# ############################
FROM alpine

RUN apk update; \
    apk upgrade; \
    apk --no-cache add bash; \
    apk --no-cache add curl tzdata;

ENV TZ=Asia/Jakarta
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

COPY --from=builder /app/server .
COPY --from=builder /app/migrations /migrations

# Install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate