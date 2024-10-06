# WORK IN PROGRESS


RabbitMQ
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.0-management

Mqtt
docker run -d --name emqx -p 18083:18083 -p 1883:1883 emqx/emqx:latest
