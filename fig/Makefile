pngs := $(patsubst %.svg,%.png,$(wildcard *.svg))

all:	${pngs}

%.png: %.svg
	inkscape --export-png=$@ $<
	convert $@ -trim output.png && mv output.png $@

clean:  
	rm -f *.png
