FROM alpine

WORKDIR /home

COPY . .

RUN chmod +x ./cloudflare-ddns-updater.sh
RUN apk update 
RUN apk add curl grep bash

RUN crontab crontab

CMD ["crond", "-f"]