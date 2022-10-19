FROM node:17-alpine

RUN apk update && apk add git zsh vim bash exa

RUN mkdir -p /var/www/app
WORKDIR /var/www/app

RUN npm install create-next-app

SHELL ["/bin/zsh", "-c"]
COPY ./.zshrc /root/.zshrc
COPY ./build-script.sh /root/build-script.sh

# ログインシェルの指定
CMD ["/bin/zsh"]
