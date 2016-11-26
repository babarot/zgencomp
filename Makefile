deps:
	go get github.com/jteeuwen/go-bindata/...

install: deps
	go-bindata -o=assets.go ./data/templates/
	go install
