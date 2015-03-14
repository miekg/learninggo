\begin{lstlisting}[numbers=none,label=src:function definition]
|\begin{tikzpicture}[overlay]
\ubrace{0.7,-1.5}{0.0,-1.5}{To declare a function, you use the \key{func} keyword \citem.}
%
\ubrace{2.1,-1.5}{0.8,-1.5}{You can optionally bind \citem{} to a specific type called %
a \first{\emph{receiver}}{receiver} (a function with a receiver is %
usually called an \index{method}{method}. This will be explored in chapter \ref{chap:interfaces}}).
%
\ubrace{3.5,-1.5}{2.4,-1.5}{Next \citem{} you write the name of your function.}
%
\ubrace{4.6,-1.5}{3.7,-1.5}{Here \citem{} we define that the variable \var{q} of type \type{int} is %
the input parameter. Parameters are passed %
\first{\emph{pass-by-value}}{pass-by-value} meaning they are copied;}
%
\ubrace{6.2,-1.5}{4.8,-1.5}{%
The variables \var{r} and \var{s} \citem{} are the %
\index{named return parameters}{named return parameters} for this function. %
Functions in Go can have multiple return values. This is very useful to return %
a value \emph{and} and error. This removes the need for %
in-band error returns (such as -1 for \texttt{EOF}) and modifying an argument.%
%
If you want the return %
parameters not to be named you only give the types: %
\lstinline{(int,int)}. If you have only one value to return you may omit %
the parentheses. If your function is a subroutine and does not have %
anything to return you may omit this entirely;}
%
\ubrace{8.5,-1.5}{6.4,-1.5}{Finally, we have the body \citem{} of the function. Note that %
\func{return} is a statement so the braces around the parameter(s) are %
optional.}
\end{tikzpicture}|
type mytype int

func (p mytype) funcname(q int) (r,s int) { return 0,0 }
||
||
\end{lstlisting}
