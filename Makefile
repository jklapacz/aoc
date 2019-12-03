CLEAN_FILES += cover.out cover.html

all: test clean

test:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

clean:
	rm -rf ${CLEAN_FILES}
