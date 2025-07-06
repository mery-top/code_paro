CC = gcc
SRC = q2/main.c \
q2/lru.c

OUT = out

$(OUT): $(SRC)
		$(CC) $(SRC) -o $(OUT)

run:
	./$(OUT)

clean:
	rm -f $(OUT)
