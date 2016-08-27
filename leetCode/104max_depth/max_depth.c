typedef struct _TreeNode
{
    int val;
    struct _TreeNode *left;
    struct _TreeNode *right;
}TreeNode;

int depth = 0;
int max_depth = 0;

void go_deep(TreeNode *root)
{
    if (NULL != root)
    {
        ++depth;
    }
    else
    {
        if (depth > max_depth)
        {
            max_depth = depth;
        }
        return;
    }
    go_deep(root->left);
    go_deep(root->right);
    --depth;
}

int maxDepth(struct TreeNode* root) {
    go_deep(root);
    int r = max_depth;
    max_depth = 0;
    depth = 0;
    return r;
}
