FROM golang:1.24.3-alpine

RUN apk add --no-cache git tzdata
ENV TZ=Asia/Jakarta

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go

COPY wait-for-it.sh /app/wait-for-it.sh

RUN chmod +x /app/wait-for-it.sh

ENTRYPOINT ["/app/wait-for-it.sh", "db:3306", "--"]

CMD ["go", "run", "./cmd/main.go"]
