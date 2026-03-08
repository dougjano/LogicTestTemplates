%:
	@go build -o bin/$*.exe ./cmd/$*
	@.\bin\$*.exe

test-%:
	@go test -v ./cmd/$*/...

.PHONY: test