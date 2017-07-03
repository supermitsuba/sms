FROM base-sms

WORKDIR /var/go/src/github.com/supermitsuba/sms/timers/time/golang

RUN  go get -d -v
RUN  go build main.go

ENV TZ=America/Detroit
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone