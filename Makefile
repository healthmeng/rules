export GOPATH=$(PWD)

bin/webui:
	go get github.com/Go-SQL-Driver/MySQL
	go install webui
test:bin/webui
	killall webui
	cd bin && ./webui &
	sleep 1
	firefox http://127.0.0.1:8888
