FROM golang:1.16-alpine
# FROM golang:1.16-buster
COPY . /go/src/api/
WORKDIR /go/src/api/
# go.modとgo.sumファイルに必要なツールを記述してインポートする
RUN apk update \
  && apk add --no-cache git alpine-sdk gnupg vim
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 /usr/bin/migrate
ENV PATH $PATH:/usr/bin/migrate
