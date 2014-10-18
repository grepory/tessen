# My Makefile-fu is terrible. This could probably be much more elegant, but
# it's a Saturday, and I have other things to do.

BINDIR := ./bin
EXAMPLES_PATH := ./examples
URL := "https://www.google.com/"

run_examples: examples
	$(BINDIR)/uptime_metric
	$(BINDIR)/http_check $(URL)

$(BINDIR):
	mkdir -p $(BINDIR)

examples: $(BINDIR) metric_examples check_examples

metric_examples: $(BINDIR) metric_uptime_example

metric_uptime_example:
	go build -o $(BINDIR)/uptime_metric $(EXAMPLES_PATH)/metric/uptime

check_examples: $(BINDIR) http_check_example

http_check_example:
	go build -o $(BINDIR)/http_check $(EXAMPLES_PATH)/check/http

.PHONY clean:
	rm -f $(BINDIR)/*
