osx:
	GOOS=darwin go build -o dwpwcheck_osx dwpwcheck.go

linux32:
	GOOS=linux GOARCH=386 go build -o dwpwcheck_linux32 dwpwcheck.go

linux64:
	GOOS=linux GOARCH=amd64 go build -o dwpwcheck_linux64 dwpwcheck.go

win32:
	GOOS=windows GOARCH=386 go build -o dwpwcheck_win32.exe dwpwcheck.go

win64:
	GOOS=windows GOARCH=amd64 go build -o dwpwcheck_win64.exe dwpwcheck.go

all: osx linux32 linux64 win32 win64