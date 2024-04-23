# Target executable name
TARGET = Lease-Agreement
SOURCE = main

$(TARGET):
	@pdflatex -output-directory=. \
			 -jobname=$(TARGET)   \
			 $(SOURCE)

# Clean up generated files
clean:
	@rm -f *.aux		 \
		   *.fdb_latexmk \
		   *.fls		 \
		   *.log		 \
		   *.out		 \
		   *.pdf		 \
		   *.synctex.gz