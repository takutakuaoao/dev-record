#
# Executes commands at the start of an interactive session.
#
# Authors:
#   Sorin Ionescu <sorin.ionescu@gmail.com>
#

# Source Prezto.
if [[ -s "${ZDOTDIR:-$HOME}/.zprezto/init.zsh" ]]; then
  source "${ZDOTDIR:-$HOME}/.zprezto/init.zsh"
fi

# Customize to your needs...

# To customize prompt, run `p10k configure` or edit ~/.p10k.zsh.
[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh

# alias
# laravel
# alias sail="bash ./vendor/bin/sail"
# alias t="sail test"

# flutter
# alias ff="flutter pub run build_runner build --delete-conflicting-outputs"
# alias ft="flutter test"

# git
alias g="git status"
alias gl="git log"
alias gp="git push origin head"
alias gc="git commit"
alias ga="git add"

# docker
# alias du="docker-compose up -d"
# alias dd="docker-compose down"
# alias de="docker exec -it"

# util
# alias l="exa -la"
alias ls="ls -la"

# GO Lang
alias got="gotest -v ./..."

# source /usr/local/lib/enhancd/init.sh
# source /opt/homebrew/opt/zsh-syntax-highlighting/share/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh

# export PATH="/opt/homebrew/opt/openjdk/bin:$PATH"
# export JAVA_HOME=$(/usr/libexec/java_home)
# export PATH="/opt/homebrew/opt/php@8.0/sbin:$PATH"
# export PATH="/opt/homebrew/opt/php@8.0/bin:$PATH"

# Go Lang
# export GOPATH=$HOME/go
# export PATH=$PATH:$GOPATH/bin
