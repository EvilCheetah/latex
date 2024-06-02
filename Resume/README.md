# Requirements
## Technology
- [LaTeX](https://www.latex-project.org/) -
    [latexmk](https://mg.readthedocs.io/latexmk.html)
- Make (**optional**)
    [[Windows](https://community.chocolatey.org/packages/make) |
    [Unix](https://www.gnu.org/software/make/)]

# Compilation
## Prerequisites
Ensure your LaTeX distribution has the following packages installed, or
install them manually:
- [`geometry`](https://ctan.org/pkg/geometry) - used for document (page)
  layout configuration.
- [`moderncv`](https://ctan.org/pkg/moderncv) - used for all the resume
  formatting in the generated document, including the header, sections, and
  section items.

For installing the packages, please refer to the following article:
[Wikibooks - LaTeX/Installing Extra Packages](https://en.wikibooks.org/wiki/LaTeX/Installing_Extra_Packages)

## Configuration
- **Visual Formatting**:  
  If you would like to configure some formatting in the document as you
  like, look at the [`config/`](./config/) directory, since the majority of
  the configuration is defined there.
- **Information (Input)**:  
  If you would like to compile the document for yourself, refer to the
  [`data/index.tex`](./data/index.tex) for the header values and to the
  [`content/`](./content/) for the body content of the resume.

## Compilation
- Using makefile, run the following command:  
  `make`

**`OR`**
- Manually:  
  `latexmk -pdf -jobname=Resume main.tex`

# Example
You can view a compiled [Resume](./Resume.pdf) filled
with placeholders to take a look at the potential result.

To see the original document that the resume was created from, you can
check
[Microsoft Word: Reference - Resume.docx](./Reference%20-%20Resume.docx)
and
[PDF: Reference - Resume.pdf](./Reference%20-%20Resume.pdf)