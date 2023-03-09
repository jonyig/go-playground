APP=go-playground

run-mail=go run main.go
dlv=--listen=:2345 --headless=true --api-version=2 --accept-multiclient


gin-html:
	  $(run-mail) ginHtml

gin-proxy:
	  $(run-mail) ginProxy

command:
	  $(run-mail) command

wire-pracetice:
	$(run-mail) wirePractice
commit-push:
	git add . && git commit -m "update" && git push origin master

make wire:
	wire ./...

dlv:
	 go build -o $(APP) -gcflags "all=-N -l"
	/Users/jonny/go/bin/dlv $(dlv) exec ./$(APP) $(s)