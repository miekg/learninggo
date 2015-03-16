\begin{Exercise}[title={Strings},difficulty=1]
\label{ex:strings}
\Question \label{ex:strings q1} Create a Go program that prints
the following (up to 100 characters):
\vskip\baselineskip
\begin{alltt}
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...
\end{alltt}
\vskip\baselineskip

\Question \label{ex:strings q2} Create a program that counts
the number of characters in this string:
\begin{alltt}
I am a G\"o programmer.
\end{alltt}
In addition, make it output the number of bytes in that string.
\emph{Hint}: Check out the \package{utf8} package.

\Question \label{ex:string q3} Extend/change the program from
the previous question to replace the three runes at
position 4 with 'abc'.

\Question \label{ex:string q4} Write a Go program
that reverses a string, so ``foobar'' is printed as ``raboof''.
\emph{Hint}: You will need to know about
conversion; skip ahead to \nref{sec:conversions} on
page \pageref{sec:conversions}.

\end{Exercise}

\begin{Answer}

\Question This program is a solution:

\begin{minipage}{0.9\textwidth}
\lstinputlisting[label=string1,numbers=none]{ex/basics/src/string1.go}
\end{minipage}

\Question To answer this question we need some help from
the \package{unicode/utf8} package. First we check the documentation
with \prog{godoc unicode/utf8 | less}. When we read the documentation
we notice \lstinline{func RuneCount(p []byte) int}. Secondly
we can convert \emph{string} to a \type{byte} slice with
\begin{lstlisting}[numbers=none]
str := "hello"
b   := []byte(str)
\end{lstlisting}
On line 2, we use a conversion (see page \pageref{sec:conversions}). Here we convert a \type{string}
to a slice of \type{byte}s. Putting this together leads to the following program.

\begin{minipage}{0.9\textwidth}
\lstinputlisting[label=src:string2,numbers=none]{ex/basics/src/string2.go}
\end{minipage}

\Question Something along the lines of:

\begin{minipage}{0.9\textwidth}
\lstinputlisting[label=src:string3,numbers=none]{ex/basics/src/string3.go}
\end{minipage}

\Question Reversing a string can be done as follows. We start from the left (\var{i}) and
the right (\var{j}) and swap the characters as we see them:

\begin{minipage}{0.9\textwidth}
\lstinputlisting[label=src:stringrev,linerange={3,},numbers=none]{ex/basics/src/stringrev.go}
\end{minipage}
\showremarks

\end{Answer}
