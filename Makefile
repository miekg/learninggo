mmark=~/g/src/github.com/miekg/mmark/mmark/mmark

all: learninggo.html

learninggo.html: *.md src/*/*.go ex/*/*.md learninggo.css head.html
	$(mmark) -page -head head.html -css learninggo.css learninggo.md > learninggo.html

clean:
	rm -f learninggo.html
