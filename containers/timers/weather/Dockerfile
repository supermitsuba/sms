FROM microsoft/dotnet:1.1.0-sdk-projectjson 

RUN apt update && apt install -y git
RUN git clone https://github.com/supermitsuba/sms.git /var/sms

WORKDIR /var/sms/timers/weather
RUN  dotnet restore
RUN  dotnet build