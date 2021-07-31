FROM golang:1.14.15

RUN mkdir -p /server
COPY . /server

WORKDIR /server
RUN go get

EXPOSE 8080

# Hot Reload
RUN [ "go", "get", "github.com/githubnemo/CompileDaemon" ]
CMD CompileDaemon -log-prefix=false -build="go build" -command="./server"