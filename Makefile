linux:
	GOOS=linux go build -o bin/linux_amd64/natter main.go

windows:
	GOOS=windows go build -o bin/windows_x86_64/natter.exe main.go

darwin:
	GOOS=darwin go build -o bin/darwin/natter main.go


all: linux windows darwin
	echo "build linux windows darwin"