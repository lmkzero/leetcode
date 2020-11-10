/*
 * @lc app=leetcode.cn id=107 lang=cpp
 *
 * [107] 二叉树的层次遍历 II
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
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int>> result;
        queue<TreeNode *> current,next;
        if(root==nullptr){
            return result;
        }else{
            current.push(root);
        }
        while(!current.empty()){
            vector<int> level;
            while(!current.empty()){
                TreeNode *node=current.front();
                current.pop();
                level.push_back(node->val);
                if(node->left!=nullptr){
                    next.push(node->left);
                }
                if(node->right!=nullptr){
                    next.push(node->right);
                }
            }
            result.push_back(level);
            swap(next,current);
        }
        reverse(result.begin(),result.end());
        return result;
    }
};
// @lc code=end

