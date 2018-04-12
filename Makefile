GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=golesson

all: clean build

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

build:
	$(GOBUILD) -o $(BINARY_NAME) -v