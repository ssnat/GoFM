$version=$args[0]

$dir="./build/v$version"

$linuxArm64Path="$dir/GoFM-linux-arm64-v$version"
$linuxAmd64Path="$dir/GoFM-linux-amd64-v$version"
$windowsArm64Path="$dir/GoFM-windows-arm64-v$version.exe"
$windowsAmd64Path="$dir/GoFM-windows-amd64-v$version.exe"

Write-Output $linuxArm64Path
$env:GOOS="linux"; $env:GOARCH="arm64"; go build -o $linuxArm64Path .

Write-Output $linuxAmd64Path
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o $linuxAmd64Path .

Write-Output $windowsArm64Path
$env:GOOS="windows"; $env:GOARCH="arm64"; go build -o $windowsArm64Path .

Write-Output $windowsAmd64Path
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o $windowsAmd64Path .

Write-Output "Done"