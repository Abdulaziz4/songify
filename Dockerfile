FROM golang:latest
COPY . ./app
WORKDIR /app
EXPOSE 8080
ENV GO111MODULE=on
ENV GOOS=linux

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download && go mod verify

COPY . .
RUN go build -o songify cmd/main.go 

EXPOSE 8080

CMD ["./songify"]
