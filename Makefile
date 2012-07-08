
dirs = 	testppack \
	arrays \
	chanasyn \
	chanfan01 \
	chanfan02 \
	chanfan03 \
	chanfan04 \
	chanfan05 \
	chanfan06 \
	chanfan07 \
	chanfan08 \
	chanloop \
	chanrange \
	chanshunt \
	chansvr1 \
	chansvr2 \
	chansyn \
	chantest \
  defer \
  envshow \
  errtest \
	fib \
	fib_embellished \
	fib_userop \
	fileread \
	flagdemo \
  force_type \
	func_lit \
	func_ret \
	gorsynchs/chanex \
	gorsynchs/mutexex \
	gorsynchs/wgroupex \
	gortns01 \
	gt044 \
	hello \
	hello_small \
	httpserv \
	interface_01 \
	interface_02 \
	interface_03 \
	interface_04 \
	interface_04a \
	interface_04b \
	list \
	list_struct \
	loghello \
	maps \
	maps_ss \
	maps_merge \
	methods \
	netcon \
	numconv \
	oltesta \
	ranges \
	rand_between \
	shax \
	showgo \
	showmac \
	showtime \
	sliceappt \
	slices \
	sltricks \
	stringlens \
	strings_utils \
	struct \
	switch_demo \
	symlshow \
	symltest \
	tcp01 \
	tcp02 \
	tcp03 \
	tcp04 \
	tcp05 \
	tcp06 \
	testgreek \
	type_String \
	unicloup \
	unicode_what \
	uniques \
	utfconv \


packlist =	numbers \

stompdirs = stomptest/receiver1 \
  stomptest/receiver2 \
  stomptest/receivernid \
  stomptest/sender \
  stomptest/sendernid \
  stomptest/sendrcv \
  stomptest/subrecv_examp \

.PHONY: $(dirs) $(packlist) packages clean format

all: $(dirs)
	@for i in $(dirs); do \
	echo $$i; \
	curd=`pwd`; \
	cd $$i && go build; \
	cd $$curd; \
	done


packages:
	@for i in $(packlist); do \
	echo $$i; \
	go install -v go-samp/$$i; \
	done

clean:
	@for i in $(dirs); do \
	echo $$i; \
	curd=`pwd`; \
	cd $$i && go clean; \
	cd $$curd; \
	done

format:
	@for i in $(dirs); do \
	echo $$i; \
	curd=`pwd`; \
	cd $$i && gofmt -w -tabwidth=2 *.go; \
	cd $$curd; \
	done

stompfmt: $(stompdirs)
	@for i in $(stompdirs); do \
	echo $$i; \
	curd=`pwd`; \
	cd $$i && gofmt -w -tabwidth=2 *.go; \
	cd $$curd; \
	done

stomptest: $(stompdirs)
	@for i in $(stompdirs); do \
	echo $$i; \
	curd=`pwd`; \
	cd $$i && go build; \
	cd $$curd; \
	done


stompclean:
	@for i in $(stompdirs); do \
	echo $$i; \
	curd=`pwd`; \
	cd $$i && go clean; \
	cd $$curd; \
	done

