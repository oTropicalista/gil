PREFIX ?= /usr
BINDIR ?= $(PREFIX)/bin
LIBDIR ?= $(PREFIX)/share/gil

build:
	@go build

install: build
	@install -v -d "$(BINDIR)/" && install -m 0755 -v gil "$(BINDIR)/gil"
	@mkdir -p "$(LIBDIR)"
	@cp libs preload.lua .gilrc.lua "$(LIBDIR)" -r
	@echo "Gil shell Installed"

uninstall:
	@rm -vrf \
			"$(BINDIR)/gil" \
			"$(LIBDIR)"
	@echo "Gil shell Uninstalled"

clean:
	@go clean

all: build install

.PHONY: install uninstall build clean