CMDDIR = ${PWD}/bin
TARGET  = $(CMDDIR)/sort
SRC = $(wildcard *.c) # all .c files
LNNAME = $(filter-out main,$(SRC:.c=)) # without main.c

all:$(TARGET)
	$(foreach name, $(LNNAME), $(shell ln -sf $(TARGET) $(CMDDIR)/$(name)))

$(TARGET):
	gcc $(SRC) -o $(TARGET)

clean:
	rm $(CMDDIR)/*

.PHONY : all clean
