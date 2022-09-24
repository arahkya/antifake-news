FROM golang:1.19

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app main.go

RUN apt -y update
RUN apt -y upgrade
RUN apt -y install sqlite3

WORKDIR /usr/src/app/Data/sqlite
RUN sqlite3 antifake-news.db < sqlite-create-table.sql

CMD ["app"]

# docker build -t antifake-news-api .
# docker run --name antifake-news-api -d -p 3001:3000 antifake-news-api
# docker run -it --rm --name antifake-news-api antifake-news-api