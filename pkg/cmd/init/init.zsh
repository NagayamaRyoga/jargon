jargon_prepare_async_callback_git() {
  JARGON_GIT_INFO="$3"
  zle reset-prompt
}

jargon_prepare_async_callback_gh() {
  JARGON_GH_INFO="$3"
  zle reset-prompt
}

jargon_prepare_async() {
  local source="$1"
  local worker="jargon_async_worker_$source"
  async_stop_worker "$worker"
  async_start_worker "$worker" -n
  async_register_callback "$worker" "jargon_prepare_async_callback_$source"
  async_job "$worker" jargon prepare --source="$source"
}

jargon_chpwd() {
  unset JARGON_GIT_INFO
  unset JARGON_GH_INFO
}

jargon_preexec() {
  unset JARGON_EXIT_STATUS_OVERWRITE
  JARGON_START="${EPOCHREALTIME}"
}

jargon_precmd() {
  JARGON_EXIT_STATUS="${JARGON_EXIT_STATUS_OVERWRITE:-$?}"
  JARGON_JOBS="${#jobstates}"
  local end="${EPOCHREALTIME}"
  JARGON_DURATION="$((${end} - ${JARGON_START:-${end}}))"
  unset JARGON_START

  jargon_prepare_async git
  jargon_prepare_async gh
}

jargon_prompt() {
  jargon prompt --exit-status="${JARGON_EXIT_STATUS}" --duration="${JARGON_DURATION}" --jobs="${JARGON_JOBS}" --width="$COLUMNS" --data-git="$JARGON_GIT_INFO" --data-gh="$JARGON_GH_INFO"
}

jargon_rprompt() {
  jargon prompt --right --exit-status="${JARGON_EXIT_STATUS}" --duration="${JARGON_DURATION}" --jobs="${JARGON_JOBS}" --width="$COLUMNS" --data-git="$JARGON_GIT_INFO" --data-gh="$JARGON_GH_INFO"
}

jargon_clear_screen() {
  JARGON_EXIT_STATUS_OVERWRITE=0
  jargon_precmd
  zle .clear-screen
}

zle -N clear-screen jargon_clear_screen

autoload -Uz add-zsh-hook
add-zsh-hook chpwd jargon_chpwd
add-zsh-hook precmd jargon_precmd
add-zsh-hook preexec jargon_preexec

setopt prompt_subst
PROMPT='$(jargon_prompt)'
RPROMPT='$(jargon_rprompt)'
