v9: v9.nn.go y.go ast.go
	go build ast.go y.go v9.nn.go

y.go: v9.y v9.nn.go
	go tool yacc v9.y

v9.nn.go: v9.nex
	nex v9.nex
