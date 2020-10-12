FROM golang:latest

LABEL maintener="FabioSebastiano <sebastiano.fabio@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 5000

RUN go build

RUN find . -name "*.go" -type f -delete

EXPOSE $PORT

CMD ["./go-gin-poc"]