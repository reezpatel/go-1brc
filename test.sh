#bin/bash

go build -o=/tmp/bin/main .

hyperfine --runs 5 '/tmp/bin/main -file=./measurements.txt'
