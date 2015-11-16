cd /var/www/appstore
killall appstore
go build
touch search.log
nohup ./appstore  >/dev/null 2>&1 &
