.DEFAULT_GOAL := help
.PHONY: help

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1\3/p' \
	| column -t  -s ' '


test_python: ## Run python tests 
	@echo "Running python tests"
	pytest 
.PHONY: test_python

run_python: ## Run python
	@echo "Running python program"
	time python3 python/euler1.py
.PHONY: run_python 

test_go: ## Run the go tests
	@echo "Running go tests for problem 1"
	go test ./go/
.PHONY: test_go

run_go: ## Run the go version to sovle problem 
	@echo "Running go program to slove problem 1"
	time go run ./go/
.PHONY: run_go

build_go:  ## Build the go program
	@echo "To build the go program"
	go build -o go_euler1 go/euler1.go


