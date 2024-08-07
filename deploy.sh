#!/bin/bash
# Copyright (c) 2024, MyOmiai Inc.
#
# This script will build bindings for OSX and Linux
#
set -e

COLOR_RESET="\033[0m"
COLOR_RED="\033[38;5;9m"
COLOR_LIGHTCYAN="\033[1;36m"
COLOR_LIGHTGREEN="\033[1;32m"

COMMANDS="aiserver client"
# shellcheck disable=SC2046
ROOT=$(cd $(dirname $0); pwd)
OUT_DIR=${OUT_DIR:-${ROOT}/bin}
BUILD_TIME=`date | sed -e 's/ /_/g'`
TARGET_OS=${TARGET_OS:-darwin}
TARGET_ARCH=${TARGET_ARCH:-arm64}

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

build() {
    local CMD=$1
    info "Building ${CMD} for ${TARGET_OS} ${TARGET_ARCH}..."

    if [[ "$TARGET_OS" = "darwin" ]] ; then
        local OUTPUT_FILE="$CMD"
    else
        local OUTPUT_FILE="${CMD}.linux"
    fi

    # ensure output dir exist
    mkdir -p ${OUT_DIR}
    local OUTPUT_PATH="${OUT_DIR}/${OUTPUT_FILE}"

    # build
    CGO_ENABLED=0 go build -ldflags "-s -w -X main.BuildTime=${BUILD_TIME} -X main.Version=${VERSION}" \
             -o $OUTPUT_PATH ${ROOT}/${CMD}.go
    if [ $? -ne 0 ] ; then
        error "Build Failed!"
    fi

    chmod +x ${OUTPUT_PATH}
}

trap '_trap' SIGINT SIGTERM

while getopts ":o:a:" flag ; do
    case $flag in
        o)
            TARGET_OS=${OPTARG}
            ;;
        a)
            TARGET_ARCH=${OPTARG}
            ;;
        *)
            echo "Usage: $0 [-a architecture - 386, amd64, arm] [-o target OS - darwin, linux]" >&2
            exit 1
            ;;
    esac
done

cd ${ROOT}
for CMD in ${COMMANDS} ; do
  echo Build ${CMD}
  GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} build $CMD
done

