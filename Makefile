.PHONY: *

#*------------VARS------------

#*______URL______

url-github = https://github.com/bastean/godo

#*______Go______

module = github.com/bastean/godo

go-tidy = go mod tidy -e

godo = go run ./cmd/godo
godo-exec = ${godo} exec -c

#*______Node______

npx = npx --no --
npm-ci = npm ci --legacy-peer-deps

release-it = ${npx} release-it -V
release-it-dry = ${npx} release-it -V -d --no-git.requireCleanWorkingDir

#*______Bash______

bash = bash -o pipefail -c

#*______Git______

git-reset-hard = git reset --hard HEAD

#*------------RULES------------

#*______Upgrade______

upgrade-managers:
	#? sudo apt update && sudo apt upgrade -y
	npm upgrade -g

upgrade-go:
	go get -t -u ./cmd/... ./internal/... ./pkg/... ./scripts/...

copydeps:
	go run ./scripts/copydeps

upgrade-node:
	${npx} ncu -ws -u
	npm i --legacy-peer-deps

upgrade-reset:
	${git-reset-hard}
	${npm-ci}

upgrade:
	${godo-exec} configs/upgrade.json

#*______Install______

install-scanners:
	curl -sSfL https://raw.githubusercontent.com/trufflesecurity/trufflehog/main/scripts/install.sh | sudo sh -s -- -b /usr/local/bin v3.63.11
	curl -sSfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sudo sh -s -- -b /usr/local/bin v0.52.2
	go install github.com/google/osv-scanner/cmd/osv-scanner@latest

install-linters:
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	npm i -g prettier

install-tools-dev: install-scanners install-linters
	go install github.com/air-verse/air@latest

install-tooling: install-tools-dev

install-tooling-ci: install-tools-dev

#*______Download______

download-dependencies:
	go mod download
	${npm-ci}

#*______Generate______

generate-required:
	go generate ./...

#*______Restore______

restore:
	${npx} husky init
	git restore .

#*______Init______

init: upgrade-managers install-tooling download-dependencies generate-required restore

init-ci: upgrade-managers install-tooling-ci download-dependencies generate-required restore

genesis:
	git init
	git add .
	$(MAKE) init

#*______ENV______

syncenv-reset:
	${git-reset-hard}

syncenv:
	cd deployments && go run ../scripts/syncenv

#*______Scan______

scan-leaks-local:
	sudo trufflehog git file://. --only-verified
	trivy repo --scanners secret .

scan-leaks-remote:
	sudo trufflehog git ${url-github} --only-verified
	trivy repo --scanners secret ${url-github}

scan-vulns-local:
	osv-scanner --call-analysis=all -r .
	trivy repo --scanners vuln .

scan-misconfigs-local:
	trivy repo --scanners misconfig .

scan-leaks: scan-leaks-local scan-leaks-remote

scan-vulns: scan-vulns-local

scan-misconfigs: scan-misconfigs-local

scans: scan-leaks scan-vulns scan-misconfigs

#*______Lint/Format______

lint:
	go mod tidy
	goimports -l -w -local ${module} .
	gofmt -l -s -w .
	${npx} prettier --ignore-unknown --write .
	$(MAKE) generate-required

lint-check:
	staticcheck ./...
	${npx} prettier --check .

#*______Test______

test-sut:
	air

test-clean: generate-required
	go clean -testcache
	mkdir -p test/report

test-unit: test-clean
	${bash} 'go test -v -cover ./pkg/... -run TestUnit.* |& tee test/report/unit.report.log'

test-integration: test-clean
	${bash} 'go test -v -cover ./pkg/... -run TestIntegration.* |& tee test/report/integration.report.log'

test-acceptance: test-clean
	${bash} 'go test -v -cover ./internal/app/... -run TestAcceptance.* |& tee test/report/acceptance.report.log'

tests: test-clean
	${bash} 'go test -v -cover ./... |& tee test/report/report.log'

#*______Build______

build: lint
	rm -rf build/
	go build -ldflags="-s -w" -o build/godo ./cmd/godo

#*______Release______

release:
	${release-it}

release-alpha:
	${release-it} --preRelease=alpha
	
release-beta:
	${release-it} --preRelease=beta

release-ci:
	${release-it} --ci --no-git.requireCleanWorkingDir $(OPTIONS)

release-dry:
	${release-it-dry}
 
release-dry-version:
	${release-it-dry} --release-version

release-dry-changelog:
	${release-it-dry} --changelog

#*______Git______

commit:
	${npx} cz

WARNING-git-forget:
	git rm -r --cached .
	git add .

WARNING-git-genesis:
	git clean -e .env* -fdx
	${git-reset-hard}
	$(MAKE) init

#*______Fix______

fix-dev: upgrade-go install-tools-dev
