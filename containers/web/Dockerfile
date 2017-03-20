FROM microsoft/dotnet:1.0.3-sdk-projectjson 

RUN apt update && apt install -y git
RUN git clone https://github.com/supermitsuba/sms.git /var/www/web

WORKDIR /var/www/web
RUN ./build.sh

WORKDIR /var/www/web/

EXPOSE 5000

ENTRYPOINT ["dotnet", "run"]