#!/bin/bash

PROGECT_DIR="$HOME/uploader"

rm -rf $PROGECT_DIR/imgs
rm -rf $PROGECT_DIR/README.md

go build $PROGECT_DIR/main.go
go build -o $PROGECT_DIR/uploader $PROJECT_DIR/main.go

echo "alias uploader=\"$PROGECT_DIR/uploader\"" >> "$HOME/.zshrc"
echo "alias uploader=\"$PROGECT_DIR/uploader\"" >> "$HOME/.bashrc"


