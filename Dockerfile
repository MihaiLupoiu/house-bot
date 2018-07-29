FROM myhay/builder:1.10-alpine as builder
MAINTAINER Mihai Lupoiu <mihai.alexandru.lupoiu@gmail.com>


# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/MihaiLupoiu/house-bot
COPY . ./
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /house-bot .

FROM alpine:latest
RUN apk update && apk upgrade
RUN apk add --no-cache ca-certificates netcat-openbsd

COPY --from=builder /house-bot ./
RUN chmod +x /house-bot

# RUN touch /config.json
# RUN touch /houses.db

ENTRYPOINT ["./house-bot"]