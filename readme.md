# Antifake News

## Containerize
1. Build App
   ```docker build -t antifake-news:latest .```
2. Create Network
   ```docker network create antifake-news```
3. Run Database
   ```docker run --name antifake-news-db -e MYSQL_ROOT_PASSWORD=vkiydKN123456 --network antifake-news -p 13306:3306 -d mysql:latest```
4. Run App
   ```docker run -p 3001:3000 --network antifake-news --name antifake-news -d antifake-news:latest```

## Note
1. [SQLite](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.3.html) Implement sample.