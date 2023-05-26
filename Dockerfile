# FROM golang:latest
# COPY . ./app
# WORKDIR /app
# EXPOSE 8080
# ENV GO111MODULE=on
# ENV GOOS=linux

# COPY go.mod go.sum ./
# RUN go mod tidy
# RUN go mod download && go mod verify

# COPY . .
# RUN go build -o songify cmd/main.go 

# EXPOSE 8080

# CMD ["./songify"]

# Build stage
FROM golang:1.20-alpine3.18 as builder
WORKDIR /app
COPY . .
RUN go build -o songify cmd/main.go 

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/songify .
COPY app.env .

EXPOSE 8080
CMD ["./songify"]


