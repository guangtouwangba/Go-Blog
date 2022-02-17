set +e
docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql
docker run -p 6379:6379 --name redis -d redis