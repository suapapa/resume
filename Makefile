ifeq ($(OS),Windows_NT)
    BROWSER = start
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        BROWSER = xdg-open
    endif
    ifeq ($(UNAME_S),Darwin)
        BROWSER = open
    endif

endif

.PHONY: all clean resume

all: resume serve

resume: *.go
	go build && ./resume

serve:
	$(BROWSER) 'http://localhost:5000'
	cd _deploy && \
	serve || (go get -v github.com/mattn/serve && serve)

deploy: resume
	rm -rf _gh-pages
	mkdir _gh-pages
	cp -r _deploy/* _gh-pages
	cd _gh-pages && git init . && git add . && \
		git commit -m "deploy" && \
		git push -f https://github.com/suapapa/resume master:gh-pages

clean:
	rm -rf resume _gh-pages _deploy/index.html
