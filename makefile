
run-mail=go run main.go

gin-html:
	  $(run-mail) ginHtml

command:
	  $(run-mail) command

wire-pracetice:
	$(run-mail) wirePractice
commit-push:
	git add . && git commit -m "update" && git push origin master

make wire:
	wire ./...