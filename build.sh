BASEDIR=$(dirname "$0")
PROGRAM="startup"

# cd to base dir
cd ${BASEDIR}

# AMD64
# linux amd64
env GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o build/Release/${PROGRAM}-linux-amd64
# macos amd64
env GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o build/Release/${PROGRAM}-macos-amd64
# windows amd64
env GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o build/Release/${PROGRAM}-win-x64.exe

# ARM64
# linux arm64
#env GOOS=linux GOARCH=arm64 go build -trimpath -ldflags "-s -w" -o build/Release/${PROGRAM}-linux-arm64
# macos arm64
#env GOOS=darwin GOARCH=arm64 go build -trimpath -ldflags "-s -w" -o build/Release/${PROGRAM}-macos-arm64

# I386
# windows i386
env GOOS=windows GOARCH=386 go build -trimpath -ldflags "-s -w" -o build/Release/${PROGRAM}-win-i386.exe