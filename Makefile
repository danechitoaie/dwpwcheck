osx:
	GOOS=darwin go build -o dwpwcheck dwpwcheck.go && zip dwpwcheck_osx.zip dwpwcheck && rm -f dwpwcheck

linux32:
	GOOS=linux GOARCH=386 go build -o dwpwcheck dwpwcheck.go && zip dwpwcheck_linux32.zip dwpwcheck && rm -f dwpwcheck

linux64:
	GOOS=linux GOARCH=amd64 go build -o dwpwcheck dwpwcheck.go && zip dwpwcheck_linux64.zip dwpwcheck && rm -f dwpwcheck

win32:
	GOOS=windows GOARCH=386 go build -o dwpwcheck.exe dwpwcheck.go && zip dwpwcheck_win32.zip dwpwcheck.exe && rm -f dwpwcheck.exe

win64:
	GOOS=windows GOARCH=amd64 go build -o dwpwcheck.exe dwpwcheck.go && zip dwpwcheck_win64.zip dwpwcheck.exe && rm -f dwpwcheck.exe

all: osx linux32 linux64 win32 win64