
dirs = arrays \
	fib \
	fib_embellished \
	fib_userop \
	hello \
	maps \
	oltesta \
	slices \
	ranges \
	testppack \

.PHONY: $(dirs)

all: $(dirs)

$(dirs):
		$(MAKE) -C $@

clean:
	@for i in $(dirs); do \
	echo $$i; \
	cd $$i && $(MAKE) clean; \
	cd ..; \
	done

