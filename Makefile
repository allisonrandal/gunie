gunie: gunie.peg.go main.go
	go build

gunie.peg.go: gunie.peg
	${GOPATH}/bin/peg -switch -inline gunie.peg

clean:
	rm -f gunie gunie.peg.go
