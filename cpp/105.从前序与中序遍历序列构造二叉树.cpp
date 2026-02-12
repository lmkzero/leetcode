/*
 * @lc app=leetcode.cn id=105 lang=cpp
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

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
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        return buildTree(begin(preorder), end(preorder), begin(inorder), end(inorder));
    }

    template <typename InputIterator>
    TreeNode* buildTree(InputIterator pre_first, InputIterator pre_last, InputIterator in_first,
                        InputIterator in_last) {
        if (pre_first == pre_last) {
            return nullptr;
        }
        if (in_first == in_last) {
            return nullptr;
        }
        auto root = new TreeNode(*pre_first);
        auto inRootPos = find(in_first, in_last, *pre_first);
        auto leftSize = distance(in_first, inRootPos);
        root->left = buildTree(next(pre_first), next(pre_first, leftSize + 1), in_first, next(in_first, leftSize));
        root->right = buildTree(next(pre_first, leftSize + 1), pre_last, next(inRootPos), in_last);
        return root;
    }
};
// @lc code=end
