pngs := $(patsubst %.svg,%.png,$(wildcard *.svg))

all:	${pngs}

%.png: %.svg
	inkscape --export-png=$@ $<

clean:  
	rm -f *.png
