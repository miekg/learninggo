all: learninggo.html

learninggo.html:
	~/g/src/github.com/mmarkdown/mmark/mmark -html -head inc/head.html -css inc/learninggo.css learninggo.md > learninggo.html

.PHONY: ast
ast:
	~/g/src/github.com/mmarkdown/mmark/mmark -ast learninggo.md

.PHONY: test
test:
	~/g/src/github.com/mmarkdown/mmark/mmark -html learninggo.md

clean:
	rm -f learninggo.html
