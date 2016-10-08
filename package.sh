# Linux
export CGO_ENABLED=0

go build -o bin/urban_linux32
go build -o bin/urban_linux64
echo "Linux build done."


# Mac
env GOOS=darwin GOARCH=386 go build -o bin/urban_darwin32
env GOOS=darwin GOARCH=amd64 go build -o bin/urban_darwin64
echo "Mac build done."

# Windows
env GOOS=windows GOARCH=386 go build -o bin/urban_win32.exe
env GOOS=windows GOARCH=amd64 go build -o bin/urban_win64.exe
echo "Windows build done."