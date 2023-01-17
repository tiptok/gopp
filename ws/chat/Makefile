
.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build
	mkdir dist
	cp chat dist/
	cp config.json dist/
	cp index.html dist/
	cp Makefile  dist/

.PHONY: clean
clean:
	rm -rf dist
	$(shell if [ -e chat ];then rm chat ;fi )
	$(shell if [ -e chat.exe ];then rm chat.exe ;fi )

.PHONY: send
send:
	scp -r dist root@106.52.103.187:~/chatsv

.PHONY: up
up:
	nohup ./chat -l=:38099 > app.log 2>&1 &
