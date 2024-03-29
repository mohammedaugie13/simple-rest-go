FROM golang:1.17.2-alpine3.14

# Install deps
RUN [ "apk", "add", "curl" ]
RUN go install github.com/cosmtrek/air@latest

# Set Workdir
WORKDIR /app
COPY  . .

RUN [ "go", "mod", "download" ]
RUN ["ls"]
EXPOSE 3000
CMD sh run.sh hot-serve