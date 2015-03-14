all: learninggo.html

learninggo.html: *.md src/*/*.go
	~/g/src/github.com/miekg/mmark/mmark/mmark learninggo.md > learninggo.html

clean:
	rm -f learninggo.html
