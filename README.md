# Calibre Bot

Этот бот был создан в связи с тем, что моя библиотека находится локально, а книги я люблю читать и в дороге, а так же [Calibre](https://calibre-ebook.com) стала достаточно медленно работать. Поэтому, была взята библиотека за основу, а так же добавлен [бот](github.com/go-telegram-bot-api/telegram-bot-api) для получения книг через Telegram. Ну и пара статей как работать с Go, чтобы собрать это вместе.


## Сборка

```
go env -w GO111MODULE=auto
export GOPATH=$PWD
go get github.com/go-telegram-bot-api/telegram-bot-api
go get github.com/mattn/go-sqlite3
go build -o bookbot src/main.go
```


## Использование

Нужно скопировать полученный bookbot в библиотеку, где находится файл metadata.db.

Далее можно запустить вручную сервис, добавив токен бота
```
./buildbot TOKEN
```
или можно оформить это в виде сервиса (Systemd Unit)

```
vim /etc/systemd/system/bookbot.service
```

```
[Unit]
Description=BookBot Service
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/PATH_TO_LIBRARY
ExecStart=/PATH_TO_LIBRARY/bookbot TOKEN

[Install]
WantedBy=multi-user.target
```
где нужно заменить PATH_TO_LIBRARY и TOKEN

После запуска в консоли или журнале вы увидите сообщение о том, что бот подключился:
```
Authorized on account ...
```
