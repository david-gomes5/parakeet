#!/bin/bash

# Set the name of your project here
PROJECT_NAME="repo-selector"
ENTRY_POINT="main.go"

# Set the path to your project if it's not in the same directory as this script
PROJECT_PATH="./"

# Set the output directory for your compiled binary
OUTPUT="build/selector"
PWD=$(pwd)

# Build the project
echo "Building $PROJECT_NAME..."
go build -o ./$OUTPUT $ENTRY_POINT

echo "Build complete: ./$OUTPUT"

# automatically add alias to your shell
if [ -f ~/.zshrc ]; then
    echo "Adding alias to ~/.zshrc..."
    echo "alias selector='$PWD/$OUTPUT'" >>~/.zshrc
else
    if [ -f ~/.bashrc ]; then
        echo "Adding alias to ~/.bashrc..."
        echo "alias selector='$PWD/$OUTPUT'" >>~/.bashrc
    else
        echo "Could not find ~/.bashrc or ~/.zshrc. Please add the alias manually."
    fi
fi