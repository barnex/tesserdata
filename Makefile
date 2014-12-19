all:
	go install
	go build tesserstamp.go
	go build tesserrate.go
	go build tessercorr.go
