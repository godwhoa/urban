# Linux
export CGO_ENABLED=0

go build -o bin/urban_linux32
go build -o bin/urban_linux64

# Mac
env GOOS=darwin GOARCH=386 go build -o bin/urban_darwin32
env GOOS=darwin GOARCH=amd64 go build -o bin/urban_darwin64

# Windows
env GOOS=windows GOARCH=386 go build -o bin/urban_win32
env GOOS=windows GOARCH=amd64 go build -o bin/urban_win64