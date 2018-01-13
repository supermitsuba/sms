FROM supermitsuba/basesms:1

# pull down git project
RUN git clone https://github.com/supermitsuba/sms.git /var/go/src/github.com/supermitsuba/sms
# install docker for running docker containers
RUN curl -sSL https://get.docker.com | sh

WORKDIR /var/go/src/github.com/supermitsuba/sms/code/golang/web

RUN  go get -d -v
RUN  go build main.go
