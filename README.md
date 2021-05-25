## Extract information from config file
---
#### Get V2ray info from config file
- Get all email
awk '/"email"/ {print $2}' /etc/xray/config.json | sed -e 's/^"//' -e 's/"$//'

- Get all uuid 
awk '/"id"/ {print $2}' /etc/xray/config.json | sed -e 's/^"//' -e 's/",$//'

- Get all username by removing @vpnje.com
awk '/"email"/ {print $2}' /etc/xray/config.json | sed -e 's/^"//' -e 's/\-.*$//' -e 's/@vpnje.com"$//'

---