# workdir info
PACKAGE=nav-server
PREFIX=$(shell pwd)
CMD_PACKAGE=${PACKAGE}
OUTPUT_DIR=${PREFIX}/bin
OUTPUT_FILE=${OUTPUT_DIR}/nav-server
COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell git describe --tags || echo "v0.0.0")
VERSION_IMPORT_PATH=${PACKAGE}/cmd
BUILD_TIME=$(shell date '+%Y-%m-%dT%H:%M:%S%Z')
VCS_BRANCH=$(shell git symbolic-ref --short -q HEAD)

# build args
BUILD_ARGS = \
    -ldflags "-X $(VERSION_IMPORT_PATH).appName=$(PACKAGE) \
    -X $(VERSION_IMPORT_PATH).version=$(VERSION) \
    -X $(VERSION_IMPORT_PATH).revision=$(COMMIT_ID) \
    -X $(VERSION_IMPORT_PATH).branch=$(VCS_BRANCH) \
    -X $(VERSION_IMPORT_PATH).buildDate=$(BUILD_TIME)"
EXTRA_BUILD_ARGS=

# which cli tools
GOLINT=$(shell which golangci-lint || echo '')
SWAG=$(shell which swag || echo '')
REDOCCLI=$(shell which redocly || echo '')

export GOCACHE=
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=

default: lint test build

lint:
	@echo "+ $@"
	@$(if $(GOLINT), , \
		$(error Please install golint: "https://golangci-lint.run/usage/install/#linux-and-windows"))
	golangci-lint run --deadline=10m -E gofmt  -E errcheck ./...

test:
	@echo "+ test"
	go test -cover $(EXTRA_BUILD_ARGS) ./...

.PHONY:build
build:
	@echo "+ build"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(BUILD_ARGS) -o /output/nav-server

dist: build
	@echo "+ $@"
	mkdir -p dist/
	@tar -cvf dist/nav-server-${VERSION}.tar README.md \
         		bin/nav-server \
         		config/config.yaml

clean:
	@echo "+ $@"
	@rm -r "${OUTPUT_DIR}"

gen-rsa-key:
	openssl genrsa -out $(PREFIX)/static/certifications/rsa_private_key.pem 2048 && \
	openssl rsa -in $(PREFIX)/static/certifications/rsa_private_key.pem \
		-pubout -out $(PREFIX)/static/certifications/rsa_public_key.pem

gen-rsa-key-pkcs8:
	openssl genrsa -out $(PREFIX)/static/certifications/keypair.pem 2048 && \
	openssl pkcs8 -topk8 -inform PEM -outform PEM -nocrypt \
		-in $(PREFIX)/static/certifications/keypair.pem \
		-out $(PREFIX)/static/certifications/pkcs8.key && \
    openssl rsa -in $(PREFIX)/static/certifications/pkcs8.key \
    	-pubout -out $(PREFIX)/static/certifications/pkcs8.pem && \
    rm -f $(PREFIX)/static/certifications/keypair.pem

# swag version >= 1.8.4
# go install github.com/swaggo/swag/cmd/swag@v1.8.4
gen-swag-admin:
	@echo "+ $@"
	@$(if $(SWAG), , \
		$(error Please install swag cli, using go: "go install github.com/swaggo/swag/cmd/swag@v1.8.4"))
	@$(if $(REDOCCLI), , \
            		$(error Please install redoc cli, using npm or yarn: "npm i -g @redocly/cli@latest"))
	swag init --dir app/admin \
 		--output app/admin/http/docs \
 		--parseDependency --parseInternal \
 		--generalInfo admin.go && \
 	redocly build-docs app/admin/http/docs/*.yaml -o app/admin/http/docs/apidoc.html && \
 	node ./redocly-copy.js app/admin/http/docs/apidoc.html



