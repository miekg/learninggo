{.exercise data-difficulty="1"}
### Pointers and reflection

One of the last paragraphs in section "\titleref{sec:introspection and reflection}"
on page \pageref{sec:introspection and reflection}, has
the following words:

> The code on the right works OK and sets the member `Name`
> to "Albert Einstein". Of course this only works when you call `Set()`
> with a pointer argument.

Why is this the case?

### Answer

When called with a non-pointer argument the variable is a copy (call-by-value).
So you are doing the reflection voodoo on a copy. And thus you are *not*
changing the original value, but only this copy.
