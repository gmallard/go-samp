
dirs = 	testppack \
	arrays \
	fib \
	fib_embellished \
	fib_userop \
	fileread \
	hello \
	interface_01 \
	interface_02 \
	maps \
	methods \
	numconv \
	oltesta \
	showgo \
	slices \
	struct \
	ranges \
	type_String \
	vector_01 \

.PHONY: $(dirs) clean gomake

all: $(dirs)

$(dirs):
		$(MAKE) -C $@

gomake:
	cd gomake; \
	gotest; \
	gomake install; \
	$(MAKE) clean; \
	cd ..

clean:
	@for i in $(dirs); do \
	echo $$i; \
	cd $$i && $(MAKE) clean; \
	cd ..; \
	done

