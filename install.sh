#!/bin/bash

PROGECT_DIR="$HOME/uploader"

cd $PROGECT_DIR
go build -o $PROJECT_DIR/main.go

rm -rf $PROGECT_DIR/.git*
rm -rf $PROGECT_DIR/imgs
rm -rf $PROGECT_DIR/README.md
rm -rf $PROGECT_DIR/main.go

echo "alias uploader=\"$PROGECT_DIR/uploader\"" >> "$HOME/.zshrc"
echo "alias uploader=\"$PROGECT_DIR/uploader\"" >> "$HOME/.bashrc"


