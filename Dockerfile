FROM golang:1.7.3
WORKDIR /go/src/github.com/zhukov174rus/Api/
RUN go get -d -v golang.org/x/net/html  
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/zhukov174rus/Api/app .
CMD ["./app"]  
#sudo docker build -t zhukov174rus/api:latest .

