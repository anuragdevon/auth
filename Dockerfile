FROM golang:alpine

WORKDIR /auth_svc

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./auth_svc ./main.go

EXPOSE 8010

CMD ./auth_svc
