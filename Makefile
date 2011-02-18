
dirs = 	testppack \
	arrays \
	chansyn \
	chanasyn \
	chansvr1 \
	chansvr2 \
	chantest \
  defer \
  envshow \
	fib \
	fib_embellished \
	fib_userop \
	fileread \
  force_type \
	func_lit \
	func_ret \
	gortns01 \
	hello \
	hello_small \
	httpserv \
	interface_01 \
	interface_02 \
	list \
	list_struct \
	loghello \
	maps \
	methods \
	multisrc \
	numconv \
	oltesta \
	ranges \
	showgo \
	slices \
	stringlens \
	strings_utils \
	struct \
	tcp01 \
	tcp02 \
	tcp03 \
	tcp04 \
	type_String \
	vector_01 \

packlist =	gomake \

.PHONY: $(dirs) $(packlist) packages clean format

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


format:
	@for i in $(dirs); do \
	echo $$i; \
	cd $$i && gofmt -w -spaces -tabwidth=2 *.go; \
	cd ..; \
	done

