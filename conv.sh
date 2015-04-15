perl -pe 's/\\emph{(.*?)}/*$1*/g' | \
perl -pe 's/\\type{(.*?)}/`$1`/g' | \
perl -pe 's/\\key{(.*?)}/`$1`/g' | \
perl -pe 's/\\func{(.*?)}/`$1`/g' | \
perl -pe 's/\\file{(.*?)}/`$1`/g' | \
perl -pe 's/\\prog{(.*?)}/`$1`/g' | \
perl -pe 's/\\mbox{(.*?)}/$1/g' | \
perl -pe 's/\\var{(.*?)}/`$1`/g' | \
perl -pe 's/\\user{(.*?)}/$1/g' | \
perl -pe 's/\\pr/\%/g' | \
perl -pe 's/\\texttt{(.*?)}/`$1`/g' | \
perl -pe 's/\\package{(.*?)}/`$1`/g' | \
perl -pe 's/\\lstinline{(.*?)}/`$1`/g' | \
perl -pe 's/\\section{(.*?)}/## $1/g' | \
perl -pe 's/\\subsection{(.*?)}/### $1/g' | \
perl -pe 's/\\cite{(.*?)}/[\@$1]/g' | \
perl -pe 's/\\label{(.*?)}//g' | \
perl -pe 's/\\ref{(.*?)}/(#$1)/g' | \
perl -pe 's/\\index{(.*?)}/((($1)))/g' | \
perl -pe 's/``/"/g' | \
perl -pe "s/''/\"/g" | \
perl -pe 's/\\noindent(\{\})?//' 
