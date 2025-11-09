FROM caddy:2.4.6-alpine

RUN mkdir -p /etc/caddy

COPY ./Caddyfile /etc/caddy
