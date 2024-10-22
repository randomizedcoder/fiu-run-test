# fiu-run-test

Quick test of fiu-run in golangs

https://blitiri.com.ar/p/libfiu/doc/posix.html

/usr/bin/fiu-run -x -c "enable_random name=posix/io/rw/read,probability=${FIO_PERCENT} name=posix/io/rw/write,probability=${FIO_PERCENT}"
