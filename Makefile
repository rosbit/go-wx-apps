SHELL=/bin/bash

EXES = wx-echo-server wx-server
TOOLS = wx-menu parseAesBody wx-userinfo wx-qr

all: .build .tools

.build:
	@for exe in $(EXES); do \
		echo "building $$exe -> samples/$$exe ..."; \
		(cd samples/$$exe; $(MAKE) -s -f ../../make.inc s=static); \
		if [ $$? -ne 0 ]; then \
			break; \
		fi; \
	done

.tools:
	@for exe in $(TOOLS); do \
		echo "building $$exe -> tools/$$exe ..."; \
		(cd tools/$$exe; $(MAKE) -s -f ../../make.inc s=static); \
		if [ $$? -ne 0 ]; then \
			break; \
		fi; \
	done

clean:
	@for exe in $(EXES); do \
		rm -f samples/$$exe/$$exe; \
	done
	@for exe in $(TOOLS); do \
		rm -f tools/$$exe/$$exe; \
	done
