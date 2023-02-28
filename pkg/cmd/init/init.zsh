jargon_chpwd() {
}

jargon_preexec() {
  JARGON_START="${EPOCHREALTIME}"
}

jargon_precmd() {
  JARGON_GIT_INFO="$(jargon prepare --source=git)"

  JARGON_EXIT_STATUS="$?"
  JARGON_JOBS="${#jobstates}"
  local end="${EPOCHREALTIME}"
  JARGON_DURATION="$((${end} - ${JARGON_START:-${end}}))"
  unset JARGON_START
}

jargon_prompt() {
  jargon prompt --exit-status="${JARGON_EXIT_STATUS}" --duration="${JARGON_DURATION}" --jobs="${JARGON_JOBS}" --width="$COLUMNS" --data-git="$JARGON_GIT_INFO"
}

jargon_rprompt() {
  jargon prompt --right --exit-status="${JARGON_EXIT_STATUS}" --duration="${JARGON_DURATION}" --jobs="${JARGON_JOBS}" --width="$COLUMNS" --data-git="$JARGON_GIT_INFO"
}

autoload -Uz add-zsh-hook
add-zsh-hook chpwd jargon_chpwd
add-zsh-hook precmd jargon_precmd
add-zsh-hook preexec jargon_preexec

setopt prompt_subst
PROMPT='$(jargon_prompt)'
RPROMPT='$(jargon_rprompt)'
