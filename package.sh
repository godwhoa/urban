# Linux
export CGO_ENABLED=0

env GOARCH=386 go build -o bin/urban_linux32
go build -o bin/urban_linux64
echo "Linux build done."

# Freebsd
env GOOS=freebsd GOARCH=386 go build -o bin/urban_freebsd32
env GOOS=freebsd GOARCH=amd64 go build -o bin/urban_freebsd64
echo "Freebsd build done."

# Openbsd
env GOOS=openbsd GOARCH=386 go build -o bin/urban_openbsd32
env GOOS=openbsd GOARCH=amd64 go build -o bin/urban_openbsd64
echo "Openbsd build done."

# Netbsd
env GOOS=netbsd GOARCH=386 go build -o bin/urban_netbsd32
env GOOS=netbsd GOARCH=amd64 go build -o bin/urban_netbsd64
echo "Netbsd build done."

# Mac
env GOOS=darwin GOARCH=386 go build -o bin/urban_darwin32
env GOOS=darwin GOARCH=amd64 go build -o bin/urban_darwin64
echo "Mac build done."

# Windows
env GOOS=windows GOARCH=386 go build -o bin/urban_win32.exe
env GOOS=windows GOARCH=amd64 go build -o bin/urban_win64.exe
echo "Windows build done."