FROM alpine:latest

WORKDIR /app

COPY listenerApp ./listenerApp

CMD ["/app/listenerApp"]