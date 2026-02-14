#!/bin/bash

PROJECT_DIR="$HOME/uploader"

go build -o "$PROJECT_DIR/uploader" "$PROJECT_DIR/main.go"

rm -rf $PROJECT_DIR/.git*
rm -rf $PROJECT_DIR/imgs
rm -rf $PROJECT_DIR/README.md
rm -rf $PROJECT_DIR/main.go

echo "alias uploader=\"$PROJECT_DIR/uploader\"" >> "$HOME/.zshrc"
echo "alias uploader=\"$PROJECT_DIR/uploader\"" >> "$HOME/.bashrc"


