#
# github.com/randomizedcoder/fiu-run-test/Makefile
#

# ldflags variables to update --version
# short commit hash
COMMIT := $(shell git describe --always)
DATE := $(shell date -u +"%Y-%m-%d-%H:%M")
BINARY := fiu-run-test

all: clean build

test:
	go test

clean:
	[ -f ${BINARY} ] && rm -rf ./${BINARY} || true

build:
	go build -ldflags \
		"-X main.commit=${COMMIT} -X main.date=${DATE} -X main.date=${VERSION}" \
		-o ./${BINARY} \
		./${BINARY}.go

run:
	FIO_PERCENT=0.5
	/usr/bin/fiu-run -x -c "enable_random name=posix/io/rw/read,probability=${FIO_PERCENT} name=posix/io/rw/write,probability=${FIO_PERCENT}" ./${BINARY}.go
