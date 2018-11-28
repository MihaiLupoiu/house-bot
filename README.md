# house-bot
```c
cd go/src/github.com/MihaiLupoiu/house-bot/
docker build -t house-bot .
docker run -d --restart unless-stopped -v ${PWD}/houses.db:/houses.db -v ${PWD}/config.json:/config.json house-bot:latest
```
