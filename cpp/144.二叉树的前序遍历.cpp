/*
 * @lc app=leetcode.cn id=144 lang=cpp
 *
 * [144] 二叉树的前序遍历
 */

#include <stack>
#include <vector>
#include "define.h"

using namespace std;

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
    vector<int> preorderTraversal(TreeNode *root) {
        vector<int> result;
        stack<TreeNode *> s;
        if (root != nullptr) {
            s.push(root);
        }
        while (!s.empty()) {
            TreeNode *p = s.top();
            s.pop();
            result.push_back(p->val);
            // 栈，先进后出
            if (p->right != nullptr) {
                s.push(p->right);
            }
            if (p->left != nullptr) {
                s.push(p->left);
            }
        }
        return result;
    }
};
// @lc code=end
