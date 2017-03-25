FROM microsoft/dotnet:1.1.0-sdk-projectjson 

RUN apt update && apt install -y git
RUN git clone https://github.com/supermitsuba/sms.git /var/www/web

WORKDIR /var/www/web/web
RUN  dotnet restore
RUN dotnet publish -c Release -o out

EXPOSE 5000

# The arguments are:
# 1 - Base url to weather service "http://api.openweathermap.org"
# 2 - Relative Path to the resource.  This has the key in it - "data/2.5/weather?zip=48307,us&appid="
# 3 - The base url for the web service to send messages to.  This is in the docker compose file.
# 4 - The relative path for the message service
CMD ["dotnet","run", "http://api.openweathermap.org", "data/2.5/weather?zip=48307,us&appid=", "web_message:8000", "api/message"]