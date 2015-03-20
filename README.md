# pctilde

pctilde is a program that emulates the behavior of [%~ for
zsh](http://stackoverflow.com/questions/13660636/what-is-percent-tilde-in-zsh).

From the zsh documentation:

```
As  %d  and %/, but if the current working directory has a named
directory as its prefix, that part is replaced by a `~' followed
by  the  name  of  the directory.  If it starts with $HOME, that
part is replaced by a `~'.
```

To use this in, eg, bash, add it to your path and then to your PS1 like this:

```bash
PS1='..$(pctilde $(pwd))..'
```

