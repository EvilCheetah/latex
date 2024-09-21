COMPILER  = latexmk
FLAGS     = -pdf
DEV_FLAGS = -pvc

# Target executable name
TARGET = "Oklahoma Secretary of State - Agreement"
SOURCE = main

$(TARGET):
	@$(COMPILER) $(FLAGS)  \
		-jobname=$(TARGET) \
		$(SOURCE).tex

dev:
	@$(COMPILER) $(FLAGS) $(DEV_FLAGS) \
		-jobname=$(TARGET)			   \
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