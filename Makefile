# Go Tools
GO  = GO111MODULE=on go
deps:
	${GO} mod vendor
	${GO} mod download

benchmark:
	${GO} test -bench Benchmark.+ -failfast -run -Benchmark.+ -v ./...

info: fmt
	depscheck -totalonly -tests .
	golocc
std-info: fmt
	depscheck -stdlib -v .
install:
	${GO} get -u github.com/divan/depscheck
	${GO} install github.com/golangci/golangci-lint/cmd/golangci-lint