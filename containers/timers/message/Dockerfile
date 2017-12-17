FROM supermitsuba/basesms:1

WORKDIR /var/go/src/github.com/supermitsuba/sms/timers/message/golang

RUN  go get -d -v
RUN  go build main.go
