#!/bin/bash

install () {
    echo "Installing sparky cli"
    git clone https://github.com/Ayobami6/sparky_generate
    chmod u+x ./bin/sparky
    sudo cp ./bin/sparky /usr/local/bin

    # clean up
    rm -rf ./sparky_generate
}

install 