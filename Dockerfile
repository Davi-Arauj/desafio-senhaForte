FROM golang:1.19

WORKDIR /go/app

RUN apt-get update

EXPOSE 8080
CMD ["tail", "-f", "/dev/null"]