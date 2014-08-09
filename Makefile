.PHONY: all binary build default test get run

ext =
ifeq "$(GOOS)" "windows"
ext =.exe
endif

osdir =
ifdef GOOS
osdir =$(GOOS)/
endif

dirs=$(wildcard 20*/*/*) 
target=$(patsubst %,bin/$(osdir)%$(ext),$(dirs))
files=$(addsuffix /main.go,$(dirs))

default: build

all: run

get:
	go get ./...

test: get 
	go test ./... -v 

run: $(files)

$(files): get
	go run $@

build: $(target)

$(target):
	go build -o $@ $(patsubst bin/$(osdir)%$(ext),%/main.go,$@) 

print-vars:
		@$(foreach v,$(.VARIABLES),$(info $v=$($v)))
