all:
	export GOARCH=amd64 && \
	echo `pwd` && \
	export LD_LIBRARY_PATH=$(shell pwd)/.libs:$$LD_LIBRARY_PATH && \
	go build

run:
	export LD_LIBRARY_PATH=$(shell pwd)/.libs:$$LD_LIBRARY_PATH && \
	./rnnoise_go_demo $(input) $(output)