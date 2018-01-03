FROM supermitsuba/basesms:1

# pull down git project
RUN git clone https://github.com/supermitsuba/sms.git /var/go/src/github.com/supermitsuba/sms

WORKDIR /var/go/src/github.com/supermitsuba/sms/serial/golang
# build golang project
RUN  go get -d -v
RUN  go build main.go
