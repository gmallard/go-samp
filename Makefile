
dirs = fib \
	fib_embellished \
	fib_userop \
	hello \
	oltesta

all:
	@for i in $(dirs); do \
	echo $$i; \
	cd $$i && $(MAKE); \
	cd ..; \
	done

clean:
	@for i in $(dirs); do \
	echo $$i; \
	cd $$i && $(MAKE) clean; \
	cd ..; \
	done

