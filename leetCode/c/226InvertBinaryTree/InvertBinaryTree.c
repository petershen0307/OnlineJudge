/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
struct TreeNode* invertTree(struct TreeNode* root) {
    if (NULL == root)
    {
        return NULL;
    }
    struct TreeNode *left;
    struct TreeNode *right;
    left = invertTree(root->left);
    right = invertTree(root->right);
    root->left = right;
    root->right = left;
    return root;
}
