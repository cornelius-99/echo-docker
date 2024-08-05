FROM golang:1.22

WORKDIR /app
COPY main.go go.mod ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /echo-test
EXPOSE 8080

CMD ["/echo-test"]