# Description
## Used Packages
- [`tabularray`](https://ctan.org/pkg/tabularray)
- [`datatool`](https://ctan.org/pkg/datatool)

## Compilation
To compile the document, you can either manually run the following command:  
`latexmk -pdf -jobname=Example main.tex`  
or use a  
[`Makefile`](https://www.gnu.org/software/make/)  
or a  
[`Taskfile`](https://taskfile.dev/)

## Error Message
When attempting to compile, the following error occurs:  
```
! Misplaced alignment tab character &.
<argument> ...annumeral \dtlforeachlevel }}\name &
                                                   \email \\ \edef \@dtl@tmp...
l.9 \end
        {tblr}
```