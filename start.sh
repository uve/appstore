cd /var/www/appstore
killall appstore
go build
nohup ./appstore  >/dev/null 2>&1 &
