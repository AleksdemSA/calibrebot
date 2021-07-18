# Calibre Bot

This telegram bot created for [Calibre](https://calibre-ebook.com)

## Run

Copy file bookbot from releases to library (where the file metadata.db is located).

You can run bot mannualy (with token!):
```
./bookbot TOKEN
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
