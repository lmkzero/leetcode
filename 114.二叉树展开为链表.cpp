/*
 * @lc app=leetcode.cn id=114 lang=cpp
 *
 * [114] 二叉树展开为链表
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    void flatten(TreeNode* root) {
        if(root==nullptr){
            return;
        }
        stack<TreeNode *> s;
        s.push(root);
        while(!s.empty()){
            auto p=s.top();
            s.pop();
            if(p->right){
                s.push(p->right);
            }
            if(p->left){
                s.push(p->left);
            }
            p->left=nullptr;
            if(!s.empty()){
                p->right=s.top();
            }
        }
    }
};
// @lc code=end

