FROM microsoft/dotnet:1.0.3-sdk-projectjson 

RUN apt update && apt install -y git
RUN git clone https://github.com/supermitsuba/sms.git /var/www/web

WORKDIR /var/www/web/web
RUN  dotnet restore
RUN dotnet publish -c Release -o out

EXPOSE 5000

CMD ["dotnet","run"]