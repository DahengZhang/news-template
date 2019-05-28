GOOS=linux go build

docker build -t dahengzhang/news .

docker run -d --name news-server --link news-mysql:db -p 8080:8080 dahengzhang/news
