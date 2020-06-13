
# VARIABLES
# -


# ENVIRONMENT VARIABLES
export GO111MODULE=on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## application

build :		## Start application
	go build

start :		## Start application
	godotenv -f local.env go run main.go

generate-swagger-doc :		## Generate Swagger documentation (both JSON and YAML formats)
	swagger generate spec -o static/swagger-ui/swagger.json
	swagger generate spec -o static/swagger-ui/swagger.yaml

open-swagger-ui :		## Open Swagger-UI
	open "http://localhost:8080/swagger-ui/index.html"

get-swagger-json :		## Get Swagger JSON
	open "http://localhost:8080/swagger-ui/swagger.json"

generate-markdown-doc-file : generate-swagger-doc		## Generate Markdown documentation from Swagger documentation
	mkdir -p docs/md/
	swagger2markup convert -i static/swagger-ui/swagger.yaml -f docs/md/swagger -c md-config.properties

generate-markdown-doc-folder : generate-swagger-doc		## Generate Markdown documentation from Swagger documentation
	mkdir -p docs/md/
	swagger2markup convert -i static/swagger-ui/swagger.yaml -d docs/md/swagger -c md-config.properties

## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
