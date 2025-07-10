# go makefile

program != basename $$(pwd)

latest_release != gh release list --json tagName --jq '.[0].tagName' | tr -d v
version != cat VERSION

gitclean = if git status --porcelain | grep '^.*$$'; then echo git status is dirty; false; else echo git status is clean; true; fi

install_dir = /usr/local/bin
postinstall =

$(program): build

build: fmt
	fix go build

fmt: go.sum
	fix go fmt . ./...

go.mod:
	go mod init

go.sum: go.mod
	go mod tidy

install: build
	doas install -m 0755 $(program) $(install_dir)/$(program) $(postinstall)

test: fmt
	go test -v -failfast . ./...

debug: fmt
	go test -v -failfast -count=1 -run $(test) . ./...

release:
	@$(gitclean) || { [ -n "$(dirty)" ] && echo "allowing dirty release"; }
	@$(if $(update),gh release delete -y v$(version),)
	gh release create v$(version) --notes "v$(version)"

clean:
	rm -f $(program)
	go clean

sterile: clean
	which $(program) && go clean -i || true
	go clean -r || true
	go clean -cache
	go clean -modcache
	rm -f go.mod go.sum
