FROM golang:1.22.2 AS build-stage
  WORKDIR /app

  COPY go.mod go.sum ./
  RUN go mod download

  COPY . .

  EXPOSE 8080

  CMD ["go", "run", "cmd/main.go"]
