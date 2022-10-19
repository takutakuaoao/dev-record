FROM node:17-alpine

RUN apk update && apk add git zsh vim bash exa

RUN mkdir -p /var/www/app
WORKDIR /var/www/app

RUN npm install create-next-app

SHELL ["/bin/zsh", "-c"]

# ENV NODE_OPTIONS=--openssl-legacy-provider

# ログインシェルの指定
CMD ["/bin/zsh"]
