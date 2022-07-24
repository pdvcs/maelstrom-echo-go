#!/bin/bash
set -e

export ARG="$1"
[[ "$ARG" == "" ]] && export ARG="build"

case "$ARG" in
    build)
        go build -o go-echo
    ;;

    deploy)
        [[ -f ./go-echo ]] && rm go-echo
        go build -o go-echo
        cp go-echo ~/repos/maelstrom-bin/    
    ;;

    clean)
        [[ -f ./go-echo ]] && rm go-echo
        true
    ;;
    
    help | --help | -h)
        echo "./build.sh [build | deploy | clean]"
        echo "'build' is the default action."
    ;;

    *)
        echo "Unknown action: $ARG (type './build.sh -h' for help)"
    ;;
esac
