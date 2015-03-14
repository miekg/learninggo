all: learninggo.html

learninggo.html: *.md src/*/*.go ex/*/*.md
	~/g/src/github.com/miekg/mmark/mmark/mmark -page learninggo.md > learninggo.html

clean:
	rm -f learninggo.html
