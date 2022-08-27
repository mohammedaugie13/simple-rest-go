#! /bin/sh

ARGS=$1

if [ "$ARGS" == "hot-serve" ]; then
	air
elif [ "$ARGS" == "serve" ]; then
    go run main.go serve $@
elif [ "$ARGS" == "setup" ]; then
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
elif [ "$ARGS" == "database" ]; then
    go run main.go database
else
    echo "Command is not found"
fi