FROM golang:1.19.0-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /github/akkishivaprakash.com

EXPOSE 8080

CMD ["/github/akkishivaprakash.com"]
