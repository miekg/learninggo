all: learninggo.html

learninggo.html:
	~/g/src/github.com/miekg/mmark/mmark/mmark learninggo.md > learninggo.html

clean:
	rm -f learninggo.html
