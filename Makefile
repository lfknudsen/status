all:
	go build
	cp _canonical_.status-codes.csv .status-codes.csv
	install
	install .status-codes.csv