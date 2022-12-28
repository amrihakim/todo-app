FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

EXPOSE 3030

ENTRYPOINT ["/app/binary"]