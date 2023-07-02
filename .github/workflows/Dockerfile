FROM golang:1.19.0-alpine
RUN mkdir /app
ADD  F:\Golang\projectgolang\cmd\main\main.exe /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main .
CMD ["F:\Golang\projectgolang\cmd\main"]
