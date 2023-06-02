FROM golang:alpine

WORKDIR /auth

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./auth ./main.go

EXPOSE 8010

CMD ./auth
