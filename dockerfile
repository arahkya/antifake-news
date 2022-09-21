FROM golang:1.19

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app main.go

CMD ["app"]

# docker build -t antifake-news-api .
# docker run -it --rm --name antifake-news-api antifake-news-api