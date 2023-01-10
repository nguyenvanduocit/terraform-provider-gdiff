default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
build:
	go build -o "./terraform-provider-gdiff_v0.0.0" main.go
