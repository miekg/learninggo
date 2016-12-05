#!/bin/bash

grep -q "/inc/prism" ./learninggo.html
HIGHLIGHTED=$?

if [ $HIGHLIGHTED -ne 0 ]; then  # Introduce syntax highlighting
  sed -i.bak '8i\
  <link rel="stylesheet" type="text/css" href="inc/prism.css">\
<script src="inc/prism.js"></script>' learninggo.html
  rm learninggo.html.bak
fi
