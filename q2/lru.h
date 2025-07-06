#pragma once
#include <fcntl.h>
#include <time.h>
#include <unistd.h>
#include <sys/stat.h>

typedef struct LRUNode{
    char path[100];
    time_t opened_at;
    off_t size;
    ino_t inode;
    struct LRUNode* next;
}LRUNode;

void create_lru(int size);
int search_lru_node(char* filepath);
void delete_lru_node();
void add_lru_node(char* filepath);
void print_lru_node();