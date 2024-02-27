#!/bin/sh

#应用的名称
AppName=sagooiot

BuildVersion=$(git describe --abbrev=0 --tags)
BuildTime=$(date +%FT%T%z)
CommitID=$(git rev-parse HEAD)

help() {
    echo "Usage: $0 [option]"
    echo "Options:"
    echo "  linux     build for Linux"
    echo "  windows   build for Windows"
    echo "  mac       build for macOS"
    echo "  macOld    build for macOS (old version)"
}


linux(){
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $AppName -a -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"
    copyFile
    cp curl.sh bin/

    cp $AppName bin/

    rm -f $AppName

}
windows(){
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $AppName.exe -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"

    copyFile

    cp $AppName.exe bin/

    rm -f $AppName.exe

}
mac(){
    go build -o $AppName -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"

    copyFile
    cp curl.sh bin/

    cp $AppName bin/

    rm -f $AppName

}

macOld(){

     GOARCH=amd64 go build -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"

    copyFile
    cp curl.sh bin/

    cp $AppName bin/

    rm -f $AppName
}

copyFile() {
    rm -rf bin
    mkdir bin
    mkdir bin/resource
    mkdir bin/resource/public
    mkdir bin/resource/public/rsa

    cp -r ./manifest/config/. bin/config/
    cp -r ./resource/rsa/. bin/resource/rsa/

    echo "${BuildVersion} $(date +%T)"
}

case "$1" in
    "linux")
        linux
        ;;
    "windows")
        windows
        ;;
    "macOld")
        macOld
        ;;
    "mac")
        mac
        ;;
    *)
        help
        ;;
esac