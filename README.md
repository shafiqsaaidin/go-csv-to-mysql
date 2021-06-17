## Extract information from config file
---
#### Get V2ray info from config file
- Get all email <br>
```bash
awk '/"email"/ {print $2}' /etc/xray/config.json | sed -e 's/^"//' -e 's/"$//'
```

- Get all uuid <br>
```bash
awk '/"id"/ {print $2}' /etc/xray/config.json | sed -e 's/^"//' -e 's/",$//'
```

- Get all username by removing @vpnje.com <br>
```bash
awk '/"email"/ {print $2}' /etc/xray/config.json | sed -e 's/^"//' -e 's/\-.*$//' -e 's/@vpnje.com"$//'
```
