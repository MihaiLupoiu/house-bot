FROM alpine:latest
MAINTAINER Mihai Lupoiu <mihai.alexandru.lupoiu@gmail.com>

ENV TELEGRAM_CHAT_ID=@telegram-chat-id
ENV TELEGRAM_BOT_ID=@telegram-bot-id

RUN apk --no-cache add ca-certificates && update-ca-certificates

COPY house-bot /bin/house-bot
RUN chmod +x /bin/house-bot

# COPY cron /var/spool/cron/crontabs/root


# CMD crond -l 2 -f

#For deploying only the image.
#FROM myhay/packt-free-learning:1