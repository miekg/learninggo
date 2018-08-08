MMARK=~/g/src/github.com/mmarkdown/mmark/mmark

all: learninggo.html

learninggo.html: *.md src/*/*.go ex/*/*.md tab/*.md inc/learninggo.css inc/head.html
	#$(MMARK) -head inc/head.html -css inc/learninggo.css learninggo.md > learninggo.html
	$(MMARK) learninggo.md > learninggo.html

clean:
	rm -f learninggo.html
