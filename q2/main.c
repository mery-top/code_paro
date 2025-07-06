#include <stdio.h>
#include "lru.h"

/*
----------------------------------------
RUN:
make
make run
from the root

OUTPUT:
Added file:q2/test/test1.txt
Added file:q2/test/test2.txt
Added file:q2/test/test3.txt
LRU File deleted: q2/test/test1.txt
File found!
Details
q2/test/test2.txt -> Size 0 | Inode 34973845 | Time: Sun Jul  6 23:18:52 2025
---------------------------------------
*/

int main(){
    create_lru(4);

    add_lru_node("q2/test/test1.txt");
    add_lru_node("q2/test/test2.txt");
    add_lru_node("q2/test/test3.txt");

    delete_lru_node();
    search_lru_node("q2/test/test2.txt");
}