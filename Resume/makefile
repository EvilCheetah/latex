# Target executable name
TARGET = Resume
SOURCE = main

$(TARGET):
	pdflatex -output-directory=. \
			 -jobname=$(TARGET)  \
			 $(SOURCE)

# Clean up generated files
clean:
	rm -f $(TARGET)