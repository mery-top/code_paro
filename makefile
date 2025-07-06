CC = gcc
SRC = q2/main.c \
q2/lru.c

OUT = q2

$(OUT): $(SRC)
		$(CC) $(SRC) -o $(OUT)

run:
	./$(OUT)

clean:
	rm -f $(OUT)
