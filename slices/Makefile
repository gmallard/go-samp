GC=8g
LD=8l
TARG=main

O_FILES=\
	slices.8

all:
	make clean
	make $(TARG)

$(TARG): $(O_FILES)
	$(LD) -o $@ $(O_FILES)
	@echo "Done. Executable is: $@"

$(O_FILES): %.8: %.go
	$(GC) -c $<

clean:
	rm -rf *.[$(OS)o] *.a [$(OS)].out _obj $(TARG) *.8

