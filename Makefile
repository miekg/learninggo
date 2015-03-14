mmark=~/g/src/github.com/miekg/mmark/mmark/mmark

all: learninggo.html

learninggo.html: *.md src/*/*.go ex/*/*.md learninggo.css head.html
	$(mmark) -page -head head.html -css learninggo.css learninggo.md > learninggo.html

learninggo.txt: *.md src/*/*.go ex/*/*.md learninggo.css head.html
	$(mmark) -xml2 learninggo.md > learninggo.xml && xml2rfc learninggo.xml

clean:
	rm -f learninggo.html
