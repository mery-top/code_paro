#include <stdio.h>
#include <stdlib.h>
#include "lru.h"
#include <string.h>
#include <fcntl.h>
#include <time.h>
#include <unistd.h>
#include <sys/stat.h>


LRUNode* head = NULL;
int max_size =0;
int curr_size= 0;

void create_lru(int size){
    max_size = size;
    curr_size = 0;
    head = NULL;
}

int search_lru_node(char* filepath){
    LRUNode* curr = head;
    while(curr){
        if(strcmp(filepath, curr->path)== 0){
            printf("File found!\n");
            printf("Details\n");
            printf("%s -> Size %lld | Inode %llu | Time: %s\n", curr->path, curr->size, curr->inode, ctime(&curr->opened_at));
            
            return 1;
        }

        curr = curr->next;
    }
    return 0;
}

void delete_lru_node(){
    if(!head) return;

    if(!head->next){
        free(head);
        head = NULL;
        return;
    }

    LRUNode* prev = NULL;
    LRUNode* curr = head;
    while(curr->next){
        prev = curr;
        curr = curr->next;
    }

    printf("LRU File deleted: %s\n",curr->path);
    prev->next = NULL;
    
    free(curr);
    curr_size--;
}

void add_lru_node(char* filepath){
    if(search_lru_node(filepath) == 1){
        printf("File already exists\n");
        return;
    }

    if(curr_size == max_size){
        delete_lru_node();
    }

    int fd = open(filepath, O_RDONLY);
    if(fd < 0){
        perror("file open error");
        return;
    }

    struct stat file_stat;
    if (fstat(fd, &file_stat) < 0) {
        perror("stat");
        close(fd);
        return;
    }

    LRUNode* newNode = (LRUNode*) malloc(sizeof(LRUNode));
    strcpy(newNode->path, filepath);
    newNode->opened_at = time(NULL);
    newNode->size = file_stat.st_size;
    newNode->inode = file_stat.st_ino;
    newNode->next = head;
    head = newNode;
    curr_size++;

    printf("Added file:%s\n", filepath);

}

void print_lru_node() {
    LRUNode *curr = head;
    printf("LRU Cache (Most → Least Used): ");
    while (curr) {
        printf("%s -> Size %lld | Inode %llu | Time: %s\n", curr->path, curr->size, curr->inode, ctime(&curr->opened_at));
        curr = curr->next;
    }
    printf("NULL\n");
}
