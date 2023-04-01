/*
 * @lc app=leetcode.cn id=129 lang=cpp
 *
 * [129] 求根到叶子节点数字之和
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
    int sumNumbers(TreeNode* root) {
        return dfs(root,0);
    }
private:
    int dfs(TreeNode *root,int sum){
        if(root==nullptr){
            return 0;
        }
        if(root->left==nullptr&&root->right==nullptr){
            return sum*10+root->val;
        }
        return dfs(root->left,sum*10+root->val)+dfs(root->right,sum*10+root->val);
    }
};
// @lc code=end

