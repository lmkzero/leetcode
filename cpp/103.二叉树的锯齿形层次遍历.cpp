/*
 * @lc app=leetcode.cn id=103 lang=cpp
 *
 * [103] 二叉树的锯齿形层次遍历
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

#include <stdlib.h>
#include <queue>
#include <vector>
#include "define.h"

using namespace std;

class Solution {
   public:
    vector<vector<int>> zigzagLevelOrder(TreeNode *root) {
        // 使用一个bool值进行记录
        vector<vector<int>> result;
        queue<TreeNode *> current, next;
        bool left_to_right = true;
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
            if (!left_to_right) {
                reverse(level.begin(), level.end());
            }
            result.push_back(level);
            left_to_right = !left_to_right;
            swap(next, current);
        }
        return result;
    }
};
// @lc code=end
