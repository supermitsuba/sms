FROM supermitsuba/base-sms:1

RUN curl -sSL https://get.docker.com | sh

WORKDIR /var/go/src/github.com/supermitsuba/sms/web/golang/

RUN  go get -d -v
RUN  go build main.go