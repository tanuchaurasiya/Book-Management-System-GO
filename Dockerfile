FROM golang:1.21.3-alpine3.17  
RUN mkdir /app 
ADD . /app 
WORKDIR /app/cmd/main  
RUN go build -o main .
CMD ["/app/cmd/main/main"]  