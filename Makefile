all:
	export GOARCH=amd64 && \
	echo `pwd` && \
	export LD_LIBRARY_PATH=$(shell pwd)/.libs:$$LD_LIBRARY_PATH && \
	go build -o rnnoise_go_demo -ldflags="-s -w"

run:
	export LD_LIBRARY_PATH=$(shell pwd)/.libs:$$LD_LIBRARY_PATH && \
	./rnnoise_go_demo $(input)