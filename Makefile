
dirs = arrays \
	fib \
	fib_embellished \
	fib_userop \
	fileread \
	hello \
	maps \
	methods \
	oltesta \
	showgo \
	slices \
	struct \
	ranges \
	testppack \
	type_String \

.PHONY: $(dirs) clean

all: $(dirs)

$(dirs):
		$(MAKE) -C $@

clean:
	@for i in $(dirs); do \
	echo $$i; \
	cd $$i && $(MAKE) clean; \
	cd ..; \
	done

