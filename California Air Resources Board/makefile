# Target executable name
COMPILER = latexmk
FLAGS    = -pdf -latexoption="--shell-escape"

TARGET = CARB-Certificate
SOURCE = main

main:
	@$(COMPILER) $(FLAGS)  \
		-jobname=$(TARGET) \
		$(SOURCE).tex

# Clean up generated files
clean:
	@rm -rf *.aux		  \
		   *.dvi		  \
		   *.fdb_latexmk  \
		   *.fls		  \
		   *.log		  \
		   *.out		  \
		   *.out.ps	      \
		   *.pdf		  \
		   *.synctex.gz	  \
		   svg-inkscape/