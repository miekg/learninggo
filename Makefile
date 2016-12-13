all: learninggo.html

learninggo.html: *.md src/*/*.go ex/*/*.md tab/*.md inc/learninggo.css inc/head.html
	mmark -page -head inc/head.html -css inc/learninggo.css learninggo.md > learninggo.html

learninggo.txt: *.md src/*/*.go ex/*/*.md tab/*.md inc/learninggo.css inc/head.html
	mmark -xml2 -page learninggo.md > learninggo.xml && xml2rfc learninggo.xml

clean:
	rm -f learninggo.html learninggo.txt
