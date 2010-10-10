GC = 8g
LD = 8l
TARG = hello

O_FILES = hello.8

all:
	make clean
	make $(TARG)

$(TARG): $(O_FILES)
	mkdir bin
	$(LD) -o bin/$@ $(O_FILES)
	@echo "Done. Executable is: bin/$@"

$(O_FILES): %.8: %.go
	$(GC) -c $<

clean:
	rm -rf *.[$(OS)o] *.a [$(OS)].out _obj $(TARG) *.8 bin

