all: learninggo.html

learninggo.html: *.md src/*/*.go ex/*/*.md tab/*.md inc/learninggo.css inc/head.html
	~/g/src/github.com/mmarkdown/mmark/mmark -html -head inc/head.html -css css/learninggo.css learninggo.md > learninggo.html

clean:
	rm -f learninggo.html
