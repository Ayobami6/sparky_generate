#!/bin/bash

install () {
    echo "Installing sparky cli"
    git clone https://github.com/Ayobami6/sparky_generate
    chmod u+x ./sparky_generate/bin/sparky
    # get the arc name
    os_name=$(uname)
    if [[ $os_name == "Darwin" ]]; then
        chmod u+x ./sparky_generate/bin/mac/sparky
        sudo cp ./sparky_generate/bin/mac/sparky /usr/local/bin
    elif [[ $os_name == "Linux" ]]; then
        sudo cp ./sparky_generate/bin/sparky /usr/local/bin
    fi

    # clean up
    rm -rf ./sparky_generate
}

install 