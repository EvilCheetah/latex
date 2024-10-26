# Target executable name
COMPILER = latexmk
FLAGS    = -pdf

TARGET = Resume - Eugene Moshchyn
SOURCE = main

$(TARGET):
	@$(COMPILER) $(FLAGS)    \
		-jobname="$(TARGET)" \
		$(SOURCE).tex

# Clean up generated files
clean:
	@rm -f *.aux		  \
		   *.dvi		  \
		   *.fdb_latexmk  \
		   *.fls		  \
		   *.log		  \
		   *.out		  \
		   *.out.ps	      \
		   *.pdf		  \
		   *.synctex.gz