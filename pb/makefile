export GO111MODULE=on
PWD := $(shell echo `pwd`)

.PHONY: echo user

all: echo user

echo:
	powerproto build -r -p $(PWD)/echo

user:
	powerproto build -r -p $(PWD)/user
