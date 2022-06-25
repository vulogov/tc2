all: compile
pre:
	go get github.com/gammazero/deque
	go get github.com/antlr/antlr4/runtime/Go/antlr
	go get github.com/sirupsen/logrus
	go get github.com/cstockton/go-conv
	go get github.com/lrita/cmap
	go get github.com/google/uuid
	go get -u -t gonum.org/v1/gonum/...
	go get -u github.com/c2fo/vfs/v6/...
	go get -u github.com/levigross/grequests
	go get -u github.com/deckarep/golang-set
	go get -u github.com/agnivade/levenshtein
	go get -u github.com/srfrog/dict
	go get -u github.com/Jeffail/gabs/v2/...
	go get github.com/google/uuid
	go get github.com/loveleshsharma/gohive
	go get -u github.com/hbollon/go-edlib
	go get -u github.com/goml/gobrain
	go get -u github.com/rocketlaunchr/dataframe-go
	go get -u github.com/stretchr/testify/assert
	go get github.com/deckarep/golang-set
	go get github.com/vulogov/tconsole
	go get github.com/pterm/pterm

build:
	rm -rf parser
	antlr4 -Dlanguage=Go  -o parser BundLexer.g4
	antlr4 -Dlanguage=Go  -o parser BundParser.g4
c:
	go build  -v .
test:
	go test -v .
rebuild: pre build dynmod c test
compile: c test
