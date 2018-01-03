FROM resin/rpi-raspbian

# install golang and set up environment
RUN apt update && apt install -y git && apt install golang
RUN mkdir /var/go
RUN mkdir /var/go/src
RUN mkdir /var/go/src/github.com
RUN mkdir /var/go/src/github.com/supermitsuba
RUN mkdir /var/go/pkg
RUN mkdir /var/go/bin
ENV GOPATH=/var/go
