#pragma once

typedef struct LRUNode{
    char* path[100];
    struct LRUNode* next;
}LRUNode;

void create_lru(int size);
int search_lru_node(char* filepath);
void delete_lru_node();
void add_lru_node(char* filepath);
