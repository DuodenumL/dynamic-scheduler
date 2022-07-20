FROM debian:stretch-slim

WORKDIR /

COPY bin/aio-scheduler /usr/local/bin

CMD ["aio-scheduler"]