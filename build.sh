#!/bin/bash

version=$1

dir="./build/v$version"

linuxArm64Path="$dir/GoFM-linux-arm64-v$version"
linuxAmd64Path="$dir/GoFM-linux-amd64-v$version"
windowsArm64Path="$dir/GoFM-windows-arm64-v$version.exe"
windowsAmd64Path="$dir/GoFM-windows-amd64-v$version.exe"

echo $linuxArm64Path
GOOS=linux GOARCH=arm64 go build -o $linuxArm64Path .

echo $linuxAmd64Path
GOOS=linux GOARCH=amd64 go build -o $linuxAmd64Path .

echo $windowsArm64Path
GOOS=windows GOARCH=arm64 go build -o $windowsArm64Path .

echo $windowsAmd64Path
GOOS=windows GOARCH=amd64 go build -o $windowsAmd64Path .

echo "Done"