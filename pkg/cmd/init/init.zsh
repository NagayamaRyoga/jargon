jargon_chpwd() {
}

jargon_preexec() {
}

jargon_precmd() {
}

jargon_prompt() {
  jargon prompt
}

jargon_rprompt() {
  jargon prompt --right
}

autoload -Uz add-zsh-hook
add-zsh-hook chpwd jargon_chpwd
add-zsh-hook precmd jargon_precmd
add-zsh-hook preexec jargon_preexec

setopt prompt_subst
PROMPT='$(jargon_prompt)'
RPROMPT='$(jargon_rprompt)'
