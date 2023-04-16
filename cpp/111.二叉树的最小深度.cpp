/*
 * @lc app=leetcode.cn id=111 lang=cpp
 *
 * [111] 二叉树的最小深度
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
    int minDepth(TreeNode* root) { return minDepth(root, false); }

   private:
    int minDepth(TreeNode* root, bool hasbrother) {
        if (!root) {
            return hasbrother ? INT_MAX : 0;
        }
        return 1 + min(minDepth(root->left, root->right != NULL), minDepth(root->right, root->left != NULL));
    }
};
// @lc code=end
