FROM FROM resin/rpi-raspbian

RUN apt update && apt install -y git && apt install golang
RUN git clone https://github.com/supermitsuba/sms.git /var/sms

WORKDIR /var/sms/serial/golang
RUN  go get -d -v
RUN  go build main.go