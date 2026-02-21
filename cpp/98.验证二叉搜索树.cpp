/*
 * @lc app=leetcode.cn id=98 lang=cpp
 *
 * [98] 验证二叉搜索树
 */

#include <stdlib.h>
#include "define.h"

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
    TreeNode* last = nullptr;
    bool isValidBST(TreeNode* root) {
        if (!root) {
            return true;
        }
        if (!isValidBST(root->left)) {
            return false;
        }
        if (last != nullptr && last->val >= root->val) {
            return false;
        }
        last = root;
        if (!isValidBST(root->right)) {
            return false;
        }
        return true;
    }
};
// @lc code=end
