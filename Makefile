
go-install-templ:
	go install github.com/a-h/templ/cmd/templ@latest

templ-generate:
	templ generate

run:
	go build && ./EMR