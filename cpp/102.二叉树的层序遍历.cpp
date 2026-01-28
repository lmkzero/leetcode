/*
 * @lc app=leetcode.cn id=102 lang=cpp
 *
 * [102] 二叉树的层序遍历
 */

#include <stdlib.h>
#include <queue>
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
    vector<vector<int>> levelOrder(TreeNode *root) {
        vector<vector<int>> result;
        queue<TreeNode *> current, next;
        if (root == nullptr) {
            return result;
        } else {
            current.push(root);
        }
        while (!current.empty()) {
            vector<int> level;
            while (!current.empty()) {
                TreeNode *node = current.front();
                current.pop();
                level.push_back(node->val);
                if (node->left != nullptr) {
                    next.push(node->left);
                }
                if (node->right != nullptr) {
                    next.push(node->right);
                }
            }
            result.push_back(level);
            swap(next, current);
        }
        return result;
    }
};
// @lc code=end
