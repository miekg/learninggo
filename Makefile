MMARK=~/g/src/github.com/mmarkdown/mmark/mmark

all: learninggo.html

learninggo.html:
	$(MMARK) -html -head inc/head.html -css inc/learninggo.css learninggo.md > learninggo.html

.PHONY: ast
ast:
	$(MMARK) -ast learninggo.md

.PHONY: test
test:
	$(MMARK) -html learninggo.md

clean:
	rm -f learninggo.html
