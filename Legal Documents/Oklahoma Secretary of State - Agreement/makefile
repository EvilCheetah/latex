COMPILER = latexmk
FLAGS    = -pdf

# Target executable name
TARGET = "Oklahoma Secretary of State - Agreement"
SOURCE = main

$(TARGET):
	@$(COMPILER) $(FLAGS)  \
		-jobname=$(TARGET) \
		$(SOURCE).tex

# Clean up generated files
clean:
	@rm -f *.aux	  \
		*.aux		  \
		*.dvi		  \
		*.fdb_latexmk \
		*.fls		  \
		*.log		  \
		*.out		  \
		*.out.ps	  \
		*.pdf		  \
		*.synctex.gz  \