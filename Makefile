
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

packlist =	gomake \

.PHONY: $(dirs) $(packlist) packages clean 

all: $(dirs)

$(dirs):
		$(MAKE) -C $@

packages:
	@for i in $(packlist); do \
	echo $$i; \
	cd $$i; \
	gomake install; \
	$(MAKE) clean; \
	cd ..; \
	done

clean:
	@for i in $(dirs); do \
	echo $$i; \
	cd $$i && $(MAKE) clean; \
	cd ..; \
	done

