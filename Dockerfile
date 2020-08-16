FROM golang:latest
ADD . /app/
WORKDIR /app/
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -a -o /komiser .

RUN chmod +x /komiser
ENTRYPOINT ["/komiser"]
EXPOSE 9092
