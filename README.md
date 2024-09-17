# Oklahoma Secretary of State - Agreement 
This is a preserved version that replicates the style of the original agreement, which can be accessed at
[OK SOS Agreement | PDF](https://www.sos.ok.gov/content/client/Agreement.pdf)

# Previews
A compiled preview of the document is available at
[Oklahoma Secretary of State - Agreement | PDF](./Oklahoma%20Secretary%20of%20State%20-%20Agreement.pdf)

# Compilation
## Requirements
- [LaTeX](https://www.latex-project.org/) -
    [latexmk](https://mg.readthedocs.io/latexmk.html)
- Make (**optional**)
    [[Windows](https://community.chocolatey.org/packages/make) |
    [Unix](https://www.gnu.org/software/make/)]

## Prerequisites
Ensure your LaTeX distribution has the following packages installed, or
install them manually:
- [`geometry`](https://ctan.org/pkg/geometry) - used for document (page) layout configuration.
- [`tcolorbox`](https://ctan.org/pkg/tcolorbox) - used for colored and framed
  text boxes.
- [`times`](https://ctan.org/pkg/times) - used for `Times New Roman` font family.
- [`titlesec`](https://ctan.org/pkg/titlesec) - used for title and subtitle
  custom formatting.

For installing the packages, please refer to the following article:
[Wikibooks - LaTeX/Installing Extra Packages](https://en.wikibooks.org/wiki/LaTeX/Installing_Extra_Packages)

## Document Compilation
- Using makefile, run the following command:  
  `make`

**`OR`**
- Manually:  
  `latexmk -pdf -jobname="Oklahoma Secretary of State - Agreement" main.tex`
