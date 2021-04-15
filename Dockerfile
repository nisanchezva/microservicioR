FROM golang:alpine

WORKDIR /go/src/github.com/nisanchezva/microservicior
COPY . .

RUN go mod tidy
RUN go build -o /go/bin/microservicior 
CMD [ "microservicior" ]