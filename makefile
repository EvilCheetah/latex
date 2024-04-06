LEASE_AGREEMENT_SRC = Lease Agreement


lease-agreement:
	latexmk -c "$(LEASE_AGREEMENT_SRC)/main.tex"