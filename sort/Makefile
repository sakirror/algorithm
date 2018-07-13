CMDDIR = ${PWD}/bin
TARGET  = $(CMDDIR)/sort
SRC = $(wildcard *.c) # all .c files
LNNAME = $(filter-out main,$(SRC:.c=)) # without main.c

all:$(TARGET)
	$(foreach name, $(LNNAME), $(shell ln -sf $(TARGET) $(CMDDIR)/$(name)))

$(TARGET): dir
	gcc $(SRC) -o $(TARGET)

clean: dir
	rm $(CMDDIR)/*

dir:
	@mkdir -p $(CMDDIR)

.PHONY : all clean dir
