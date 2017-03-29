FROM resin/rpi-raspbian

RUN apt update && apt install -y git && apt install golang
RUN mkdir /var/go
RUN mkdir /var/go/src
RUN mkdir /var/go/src/github.com
RUN mkdir /var/go/src/github.com/supermitsuba
RUN mkdir /var/go/pkg
RUN mkdir /var/go/bin
ENV GOPATH=/var/go
RUN git clone https://github.com/supermitsuba/sms.git /var/go/src/github.com/supermitsuba/sms

WORKDIR /var/go/src/github.com/supermitsuba/sms/timers/message/golang

RUN  go get -d -v
RUN  go build main.go
