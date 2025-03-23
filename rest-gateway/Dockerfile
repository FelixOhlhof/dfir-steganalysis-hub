FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o restgw

CMD ["./restgw"]