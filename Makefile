RELEASE = 1.0
PRODPATH = build
APPNAME = packtd

# Source
objects = main.go

all: lamd64 l386 damd64 d386 wamd64 w386

lamd64: $(objects)
	GOOS=linux GOARCH=amd64 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_linux-amd64 $(objects)
l386: $(objects)
	GOOS=linux GOARCH=386 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_linux-386 $(objects)
damd64: $(objects)
	GOOS=darwin GOARCH=amd64 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_darwin-amd64 $(objects)
d386: $(objects)
	GOOS=darwin GOARCH=386 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_darwin-386 $(objects)
wamd64: $(objects)
	GOOS=windows GOARCH=amd64 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_windows-amd64 $(objects)
w386: $(objects)
	GOOS=windows GOARCH=386 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_windows-386 $(objects)