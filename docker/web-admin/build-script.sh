#!/bin/zsh

git clone --recursive https://github.com/sorin-ionescu/prezto.git ${ZDOTDIR:-$HOME}/.zprezto

setopt EXTENDED_GLOB

for rcfile in ${ZDOTDIR:-$HOME}/.zprezto/runcoms/^README.md(.N); do
    ln -s $rcfile ${ZDOTDIR:-$HOME}/.${rcfile:t}
done

# Add zstyle :prezto:module:prompt theme powerlevel10k to ~/.zpreztorc.

# fzf
git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
~/.fzf/install

# enhancd
git clone https://github.com/b4b4r07/enhancd.git ~/enhancd
最終行に.zshrcに source ~/enhancd/init.sh を追加
