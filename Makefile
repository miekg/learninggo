MMARK=~/g/src/github.com/mmarkdown/mmark/mmark

all: learninggo.html

.PHONY: learninggo.html
learninggo.html:
	$(MMARK) -html -head inc/head.html -css inc/learninggo.css learninggo.md > learninggo.html

.PHONY: learninggo.xml
learninggo.xml:
	$(MMARK) learninggo.md > learninggo.xml

.PHONY: learninggo.txt
learninggo.txt: learninggo.xml
	xml2rfc --v3 learninggo.xml

.PHONY: ast
ast:
	$(MMARK) -ast learninggo.md

.PHONY: test
test:
	$(MMARK) -html learninggo.md

clean:
	rm -f learninggo.html
