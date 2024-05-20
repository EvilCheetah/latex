# Requirements
## Technology
- [LaTeX](https://www.latex-project.org/) - [latexmk](https://mg.readthedocs.io/latexmk.html)
- Make (**optional**) [[Windows](https://community.chocolatey.org/packages/make) | [Unix](https://www.gnu.org/software/make/)]

# Compilation
## Prerequisites
Ensure your LaTeX distribution has the following packages installed, or install them manually:
- [`amsmath`](https://ctan.org/pkg/amsmath) - used for `\underset` and `\overset`, specifically for Ink Fillable Field with a `(if applicable)` Subscript on the last page.
- [`bookmark`](https://ctan.org/pkg/bookmark) - used for fixing a `rerunfilecheck` warning.
- [`enumitem`](https://ctan.org/pkg/enumitem) - used for custom reference labels in `enumeration`.
- [`etoolbox`](https://ctan.org/pkg/etoolbox) - used for conditional rendering.
- [`fancyhdr`](https://ctan.org/pkg/fancyhdr) - used for header and footer custom formatting.
- [`fontenc`](https://ctan.org/pkg/fontenc) - used for font encoding, which adds additional characters.
- [`footmisc`](https://ctan.org/pkg/footmisc) - used for footnote configuration and formatting.
- [`geometry`](https://ctan.org/pkg/geometry) - used for document (page) layout configuration.
- [`hyperref`](https://ctan.org/pkg/hyperref) - used for hyperlinks and navigation in the document.
- [`soul`](https://ctan.org/pkg/soul) - used for additional text styles.
- [`tabularx`](https://ctan.org/pkg/tabularx) - used for adjustable-width columns in tables.
- [`titlesec`](https://ctan.org/pkg/titlesec) - used for title and subtitle custom formatting.
- [`xhfill`](https://ctan.org/pkg/xhfill) - used for flexible underline for input fields in the text.
- [`zref-totpages`](https://ctan.org/pkg/zref-totpages) - used for last page reference, specifically for `Page {PAGE_NUMBER} of {TOTAL_PAGES}`.

For installing the packages, please refer to the following article: [Wikibooks - LaTeX/Installing Extra Packages](https://en.wikibooks.org/wiki/LaTeX/Installing_Extra_Packages)

## Configuration
- **Visual Formatting**:  
  If you would like to configure some formatting in the document as you like, look at the [`config/`](./config/) directory, since the majority of the configuration is defined there.
- **Information (Input)**:  
  If you would like to compile the document for yourself, refer to the [`data/index.tex`](./data/index.tex) for the input information and [`data/fallback.tex`](./data/fallback.tex) for fallback values of trailer and tractor.

## Compilation
- Using makefile, run the following command:  
  `make`

**`OR`**
- Manually:  
  `latexmk -pdf -jobname=Lease-Agreement main.tex`

# Example
You can view a compiled [Lease Agreement](./Lease-Agreement.pdf) filled with placeholders to take a look at the potential result.

# Features
The following document conditionally renders the tractor and trailer based on the input information. If you want to include `tractor` or `trailer` in the resulting document, you must ensure you have the following commands:
- For `Tractor`:
  * `\VehicleYear` - Tractor's Year
  * `\VehicleMake` - Tractor's Make
  * `\VehicleModel` - Tractor's Model
  * `\VehicleVIN` - Tractor's VIN
- For `Trailer`:
  * `\TrailerYear` - Trailer's Year
  * `\TrailerMake` - Trailer's Make
  * `\TrailerVIN` - Trailer's VIN

If you would like to customize the behavior of conditional rendering, refer to the following files:
- For `Tractor` - [content/preamble/equipment/tractor.tex](./content/preamble/equipment/tractor.tex)
- For `Trailer` - [content/preamble/equipment/trailer.tex](./content/preamble/equipment/trailer.tex)
