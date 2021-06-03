# Calibre Bot

This telegram bot creted for [Calibre](https://calibre-ebook.com)

## Build

```
go env -w GO111MODULE=auto
export GOPATH=$PWD
go get github.com/go-telegram-bot-api/telegram-bot-api
go get github.com/mattn/go-sqlite3
go build -o bookbot src/main.go
```


## Run

Copy file bookbot in library (where the file metadata.db is located).

You can run bot mannualy (with token!):
```
./buildbot TOKEN
```

or use system Systemd Unit

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
Restart=always
WorkingDirectory=/PATH_TO_LIBRARY
ExecStart=/PATH_TO_LIBRARY/bookbot TOKEN

[Install]
WantedBy=multi-user.target
```
replace PATH_TO_LIBRARY and TOKEN

In journal (journalctl -f -u bookbot) you can see this:
```
Authorized on account ...
```
