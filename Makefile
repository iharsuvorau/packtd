RELEASE = 1.1
PRODPATH = build
APPNAME = packtd

# Source
objects = main.go

all: lamd64 l386 damd64 d386 wamd64 w386

lamd64: $(objects)
	GOOS=linux GOARCH=amd64 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_linux-amd64/$(APPNAME) $(objects)
l386: $(objects)
	GOOS=linux GOARCH=386 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_linux-386/$(APPNAME) $(objects)
damd64: $(objects)
	GOOS=darwin GOARCH=amd64 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_darwin-amd64/$(APPNAME) $(objects)
d386: $(objects)
	GOOS=darwin GOARCH=386 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_darwin-386/$(APPNAME) $(objects)
wamd64: $(objects)
	GOOS=windows GOARCH=amd64 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_windows-amd64/$(APPNAME).exe $(objects)
w386: $(objects)
	GOOS=windows GOARCH=386 go build -o $(PRODPATH)/$(APPNAME)_$(RELEASE)_windows-386/$(APPNAME).exe $(objects)