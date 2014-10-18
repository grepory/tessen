# Tessen

Tessen is a modular framework for Sensu plugins written in Go. The aim is to
decouple Sensu checks, metrics, handlers, etc. from Ruby in an effort to make
them more portable and more easily distributable.

## Getting Started

Clone this repository!

```
$ git clone https://github.com/grepory/tessen && cd tessen
```

Running `make` will then build all of the examples and then run them:

```
$ make
go build -o ./bin/uptime_metric ./examples/metric/uptime
go build -o ./bin/http_check ./examples/check/http
./bin/uptime_metric
gregs-mbp	1413585004.000000	1413659359

./bin/http_check "https://www.google.com/"
HTTPCheck OK: 200 - Found - https://www.google.com/
```

## Development

Tessen is currently under somewhat active development. I don't _expect_ the
interfaces/types to change in the check or metric modules at this point, but
their functionality may somewhat. There may also be undiscovered bugs, as I
haven't written any tests yet. Feel free to experiment, however, and write
checks and metrics collectors. Do not hestitate to contact me on Twitter or
Freenode if you need anything (grepory on both).
