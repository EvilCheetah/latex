COMPILER = latexmk
FLAGS    = -pdf

# Target executable name
TARGET = Lease-Agreement
SOURCE = main

$(TARGET):
	@$(COMPILER) $(FLAGS)  \
		-jobname=$(TARGET) \
		$(SOURCE).tex

# Clean up generated files
clean:
	@rm -f *.aux		 \
		   *.fdb_latexmk \
		   *.fls		 \
		   *.log		 \
		   *.out		 \
		   *.pdf		 \
		   *.synctex.gz