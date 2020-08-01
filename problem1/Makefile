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
