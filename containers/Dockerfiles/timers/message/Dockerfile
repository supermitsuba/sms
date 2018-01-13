FROM supermitsuba/basesms:1

RUN git clone https://github.com/supermitsuba/sms.git /var/go/src/github.com/supermitsuba/sms
WORKDIR /var/go/src/github.com/supermitsuba/sms/code/golang/timers/message

RUN  go get -d -v
RUN  go build main.go
