#!/bin/bash
# Copyright (c) 2024, MyOmiai Inc.
#
# This script will build bindings for OSX and Linux
#

COLOR_RESET="\033[0m"
COLOR_RED="\033[38;5;9m"
COLOR_LIGHTCYAN="\033[1;36m"
COLOR_LIGHTGREEN="\033[1;32m"

COMMANDS="aiserver client"
# shellcheck disable=SC2046
ROOT=$(cd $(dirname $0); pwd)
BINDIR=${BINDIR:-${ROOT}/bin}
AI_USER="myomiaiadmin"
AI_PASSWORD=${AI_PASSWORD:-specify_ai_password_in_env}
AI_SERVER="ai.mo"
DESTDIR=${DESTDIR:-/home/${AI_USER}/myomiai_servers/aiserver}
DEST=${AI_USER}@${AI_SERVER}:${DESTDIR}
SSHPASS=/opt/homebrew/bin/sshpass

error() {
    echo -e "${COLOR_RED}ERROR: $1${COLOR_RESET}" >&2
    exit 1
}

warn() {
    echo -e "${COLOR_RED}WARNING: $1${COLOR_RESET}"
}

info() {
    echo -e "${COLOR_LIGHTCYAN}$1${COLOR_RESET}"
}

success() {
    echo -e "${COLOR_LIGHTGREEN}$1${COLOR_RESET}"
}

_trap() {
  echo interrupted >&2
  exit 1
}

deploy() {
    local CMD=$1
    local source=${BINDIR}/${CMD}.linux
    local dest=${DEST}/${CMD}
    info "Deploying ${source} to ${dest}..."
    ${SSHPASS} -p ${AI_PASSWORD} scp ${source} ${dest}
}

trap '_trap' SIGINT SIGTERM

cd ${ROOT}
for CMD in ${COMMANDS} ; do
  deploy ${CMD}
done

