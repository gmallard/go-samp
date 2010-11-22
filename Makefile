
dirs = 	testppack \
	arrays \
	chansyn \
	chanasyn \
	fib \
	fib_embellished \
	fib_userop \
	fileread \
	func_ret \
	gortns01 \
	hello \
	hello_small \
	httpserv \
	interface_01 \
	interface_02 \
	maps \
	methods \
	multisrc \
	numconv \
	oltesta \
	ranges \
	showgo \
	slices \
	struct \
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

