FROM supermitsuba/basesms:1

RUN git clone https://github.com/supermitsuba/sms.git /var/go/src/github.com/supermitsuba/sms
WORKDIR /var/go/src/github.com/supermitsuba/sms/code/golang/timers/time

RUN  go get -d -v
RUN  go build main.go

ENV TZ=America/Detroit
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
