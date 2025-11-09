FROM alpine:latest
WORKDIR /app
COPY frontApp /app/frontApp
CMD ["/app/frontApp"]