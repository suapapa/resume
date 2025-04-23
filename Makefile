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
	$(BROWSER) 'http://localhost:5001'
	cd _deploy && \
	serve || (go install github.com/mattn/serve@latest && serve -a :5001)

deploy: resume
	rm -rf _gh-pages
	mkdir _gh-pages
	cp -r _deploy/* _gh-pages
	cd _gh-pages && git init . && git add . && \
		git commit -m "deploy" && \
		git push -f git@github.com:suapapa/resume.git main:gh-pages

clean:
	rm -rf resume _gh-pages _deploy/index.html
