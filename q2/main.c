#include <stdio.h>
#include "lru.h"


int main(){
    create_lru(4);

    for(int i=0; i<4; i++){
        add_lru_node("test\n");
    }

    delete_lru_node();
    print_lru_node();
}