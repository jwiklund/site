### Setup

### DNS through zzzz.io

`/etc/zzzz.io` :

```bash
#!/usr/bin/env bash
FILE=/tmp/zzzz.io
URL=https://zzzz.io/api/v1/update/$HOST/?token=$TOKEN
date >> $FILE
curl -s "$URL" >> $FILE
echo >> $FILE
```

`crontab -e` :

```bash
0 5 * * * /etc/zzzz.io # update zzzz.io
```

### SSL through Lets Encrypt [acmwrapper](https://github.com/dkumor/acmewrapper) 

`site.service` :

```properties
[Unit]
Description=site

[Service]
User=USER
Environment=SSL_HOST=born.zzzz.io
WorkingDirectory=/home/USER
ExecStart=/home/USER/site

[Install]
WantedBy=multi-user.target
```

Alt self signed for testing [example](https://gist.github.com/denji/12b3a568f092ab951456).

### SystemD setup

Can not enable systemd service using link, first copy file to enable, then replace with link

```bash
sudo cp /home/USER/site.service /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl enable site
sudo rm /etc/systemd/system/site.service
sudo ln -s /home/USER/site.service /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl status site
```

### Allow listen on low ports

`sudo setcap cap_net_bind_service=+ep site` 