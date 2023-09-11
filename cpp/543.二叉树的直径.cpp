/*
 * @lc app=leetcode.cn id=543 lang=cpp
 *
 * [543] 二叉树的直径
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

#include <algorithm>
#include <iostream>
#include "define.h"

using namespace std;

class Solution {
public:
    int diameterOfBinaryTree(TreeNode* root) {
        maxDiameter = 0;
        if (root == nullptr) {
            return maxDiameter;
        }
        dfs(root);
        return maxDiameter;
    }
    int dfs(TreeNode *root) {
        if (root == nullptr) {
            return 0;
        }
        int left = dfs(root->left);
        int right = dfs(root->right);
        maxDiameter = max(maxDiameter, left+right);
        return max(left,right)+1;
    }
private:
    int maxDiameter;
};
// @lc code=end

