docker stop goblog-db
docker rm goblog-db
docker stop goblog
docker rm goblog
docker rmi goblog
make image
docker-compose up