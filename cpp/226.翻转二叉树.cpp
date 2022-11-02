/*
 * @lc app=leetcode.cn id=226 lang=cpp
 *
 * [226] 翻转二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        if(root==nullptr){
            return root;
        }
        swap(root->left,root->right);
        invertTree(root->left);
        invertTree(root->right);
        return root;
    }

    // 迭代法
    // TreeNode* invertTree(TreeNode* root) {
    // if (root == NULL) return root;
    // stack<TreeNode*> st;
    // st.push(root);
    // while(!st.empty()) {
    //     TreeNode* node = st.top();
    //     st.pop();
    //     swap(node->left, node->right);
    //     if(node->left) st.push(node->left);
    //     if(node->right) st.push(node->right);
    // }
};
// @lc code=end

