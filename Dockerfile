# Start from the latest golang base image
FROM golang:latest

LABEL maintainer="karuppaiah.al@gmail.com"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o counter .

EXPOSE 8080

# Command to run the executable
CMD ["./counter"]