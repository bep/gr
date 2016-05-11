DIRS=$(shell go list -f {{.Dir}} ./...)

check: fmt vet lint test

cyclo:
	@for d in $(DIRS) ; do \
		if [ "`gocyclo -over 20 $$d | tee /dev/stderr`" ]; then \
			echo "^ cyclomatic complexity exceeds 20, refactor the code!" && echo && exit 1; \
		fi \
	done

fmt:
	@for d in $(DIRS) ; do \
		if [ "`gofmt -l -s $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ improperly formatted go files" && echo && exit 1; \
		fi \
	done

lint:
	@if [ "`golint ./... | tee /dev/stderr`" ]; then \
		echo "^ golint errors!" && echo && exit 1; \
	fi

test:
	gopherjs test github.com/bep/gr/tests

vet:
	@if [ "`go vet ./... | tee /dev/stderr`" ]; then \
		echo "^ go vet errors!" && echo && exit 1; \
	fi

