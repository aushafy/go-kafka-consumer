FROM golang:latest

WORKDIR /app

COPY go.mod /app

RUN go mod tidy && go mod vendor

COPY . /app

CMD [ "go", "run", "/app/main.go" ]