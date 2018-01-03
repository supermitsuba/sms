# containers

This section describes how to build containers for the sms project.  I have provided these images on docker hub:

1.  [Base image](https://hub.docker.com/r/supermitsuba/basesms/)
  * This is the base image for all the images for this project.  Contains setup for golang.
2.  [Website and API](https://hub.docker.com/r/supermitsuba/web_message/)
  * Contains the management website and API for making messages.
3.  [Service to Process Messages](https://hub.docker.com/r/supermitsuba/client_sms/)
  * Used to process messages and display them on the LED
4.  [Rabbit MQ](https://hub.docker.com/r/supermitsuba/rpi-rabbitmq/)
  * Message queue to display on LED sign
5.  [Nginx](https://hub.docker.com/r/supermitsuba/nginx/)
  * Used to smooth over the API and website urls
6.  [Timer - Message](https://hub.docker.com/r/supermitsuba/message_timer/)
  * Application to send a message to web message API
7.  [Timer - Forecast](https://hub.docker.com/r/supermitsuba/forecast_timer/)
  * Application to send forecast to web message API (5 day forecast)
8.  [Timer - Time](https://hub.docker.com/r/supermitsuba/time_timer/)
  * Application to send time to the web message API
9.  [Timer - Weather](https://hub.docker.com/r/supermitsuba/weather_timer/)
  *  Application to send current weather to the web message API

## Commands to help with containers

### docker-compose
1. ```docker-compose up -d ```
  * start all images in docker-compose

### docker
1.  ```docker push supermitsuba/client_sms:1 ```
  * push a docker image to docker hub
2.  ```docker build -t supermitsuba/client_sms:1 . --no-cache ```
  * build a docker image with a tag

### Cron jobs
1. ``` */12 6-23 * * * docker run containers_message_timer ./main http://192.168.10.115:8000/api/message "message" ```
  * Display message on timer
2. ```4,9,14,19,24,29,34,39,44,49,54,59 6-23 * * * docker start time```
  * Display time on timer
3. ```*/15 6-23 * * * docker start weather```
  * Display weather on timer
4. ```*/15 6-23 * * * docker start forecast```
  * Display forecast on timer

### nginx configuration
1. [nginx configuration](https://raw.githubusercontent.com/supermitsuba/sms/master/containers/nginx.conf)


## Architecture

![](https://raw.githubusercontent.com/supermitsuba/sms/master/architecture.png)