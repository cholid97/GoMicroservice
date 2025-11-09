FROM alpine:latest

WORKDIR /app

COPY mailerApp ./mailerApp
COPY cmd/template/ ./template/

CMD ["/app/mailerApp"]