#!/bin/bash

PROGECT_DIR="$HOME/uploader"

rm -rf $PROGECT_DIR/imgs
rm -rf $PROGECT_DIR/README.md

echo "alias uploader=\"go run $PROGECT_DIR/main.go\"" >> "$HOME/.zshrc"
echo "alias uploader=\"go run $PROGECT_DIR/main.go\"" >> "$HOME/.bashrc"

