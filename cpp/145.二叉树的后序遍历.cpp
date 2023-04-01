/*
 * @lc app=leetcode.cn id=145 lang=cpp
 *
 * [145] 二叉树的后序遍历
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
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> result;
        stack<TreeNode *> s;
        TreeNode *p=root;
        TreeNode *q=nullptr;
        do{
            while(p!=nullptr){
                s.push(p);
                p=p->left;
            }
            q=nullptr;
            while(!s.empty()){
                p=s.top();
                s.pop();
                if(p->right==q){
                    result.push_back(p->val);
                    q=p;
                }else{
                    s.push(p);
                    p=p->right;
                    break;
                }
            }
        }while(!s.empty());
        return result;
    }
};
// @lc code=end

