all:
	go install
	go build tesserstamp.go
	go build tesserrate.go
	go build tesserhist.go
	go build tessercorr.go
	go build tesserlife.go
