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

Architecture
============
![](https://raw.githubusercontent.com/supermitsuba/sms/master/architecture.png)
