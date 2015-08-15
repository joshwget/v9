v9: v9.nn.go y.go ast.go variable.go
	go build -o v9 y.go v9.nn.go ast.go variable.go

y.go: v9.y v9.nn.go
	go tool yacc v9.y

v9.nn.go: v9.nex
	nex v9.nex

clean:
	rm -f v9 v9.nn.go y.output y.go
