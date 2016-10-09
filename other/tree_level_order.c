#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

typedef struct _TreeNode{
    int val;
    struct _TreeNode *left;
    struct _TreeNode *right;
}TreeNode;

typedef struct _TreeNodeList{
    TreeNode *val;
    struct _TreeNodeList *next;
}TreeNodeList;

#define MAX_LEVEL 10 //2^10-1
TreeNodeList *level_node[MAX_LEVEL];

int tree_level = -1;
void tree_level_order(TreeNode *root){
    if (NULL == root){
        return;
    }
    ++tree_level;
    assert(MAX_LEVEL > tree_level);
    TreeNodeList *new_node = malloc(sizeof(TreeNodeList));
    new_node->val = root;
// for the list order left to right
    new_node->next = level_node[tree_level];
    level_node[tree_level] = new_node;

// for the list order left to right
    tree_level_order(root->right);
    tree_level_order(root->left);
    --tree_level;
}

void free_level_nodes(){
    for (int i = 0; i < MAX_LEVEL; ++i){
        TreeNodeList *tmp = level_node[i];
        while (NULL != tmp) {
            TreeNodeList *del = tmp;
            tmp = tmp->next;
            free(del);
        }
    }
}

// 有著同樣的for loop，感覺可以用yield抽出一個for loop function，外層再包一層實作

void print_level_order(){
    for (int i = 0; i < MAX_LEVEL; ++i){
        TreeNodeList *tmp = level_node[i];
        while (NULL != tmp) {
            printf("level:%d, val:%d\n", i, tmp->val->val);
            tmp = tmp->next;
        }
    }
}

int main(){
    int a = 10;
    memset(level_node, 0, MAX_LEVEL);
    TreeNode root = {.val = 0, .left = NULL, .right = NULL};
    TreeNode node1 = {.val = 1, .left = NULL, .right = NULL};
    TreeNode node2 = {.val = 2, .left = NULL, .right = NULL};
    TreeNode node3 = {.val = 3, .left = NULL, .right = NULL};
    TreeNode node4 = {.val = 4, .left = NULL, .right = NULL};
    TreeNode node5 = {.val = 5, .left = NULL, .right = NULL};
    TreeNode node6 = {.val = 6, .left = NULL, .right = NULL};
    root.left = &node1;
    root.right = &node2;
    node1.left = &node3;
    node1.right = &node4;
    node2.left = &node5;
    node2.right = &node6;
    tree_level_order(&root);
    print_level_order();
    //free_level_nodes();
}
