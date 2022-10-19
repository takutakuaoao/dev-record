FROM golang:1.18-alpine3.15

RUN apk update && apk add git zsh vim bash exa
RUN apk add --no-cache alpine-sdk build-base

RUN mkdir -p /var/www/cli
RUN mkdir /usr/share/fonts

# シェルの設定
SHELL ["/bin/zsh", "-c"]
COPY ./docker/.zshrc /root/.zshrc
COPY ./docker/build-script.sh /root/build-script.sh

# フォント
COPY ./docker/fonts /usr/share/fonts

# Goの設定
WORKDIR /var/www/cli
COPY ./cli/go.mod /var/www/cli
COPY ./cli/go.sum /var/www/cli
RUN go mod download

COPY ./cli /var/www/cli

# ログインシェルの指定
CMD ["/bin/zsh"]
