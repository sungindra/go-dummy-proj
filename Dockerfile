FROM golang:1.21.6-bookworm

WORKDIR /app

COPY . .

RUN make vendor

ENTRYPOINT ["go", "run", "main.go"]
