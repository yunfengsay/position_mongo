git pull
kill -9 $(ps -ef | grep "28018" | grep -v grep | awk '{print $2}')
kill -9 $(ps -ef | grep "positiongo" | grep -v grep | awk '{print $2}')
kill -9 $(ps -ef | grep "caddy" | grep -v grep | awk '{print $2}')

sudo mongod    --dbpath=/usr/local/var/mongodb  --port 28018 &
go run main.go positiongo &
caddy &