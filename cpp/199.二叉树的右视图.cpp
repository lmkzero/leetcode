/*
 * @lc app=leetcode.cn id=199 lang=cpp
 *
 * [199] 二叉树的右视图
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
    vector<int> rightSideView(TreeNode* root) {
        if (!root) {
            return {};
        }
        // bfs 层序遍历 将每层最后一个加入结果数组
        vector<int> ans;
        queue<TreeNode*> que;
        que.push(root);
        while (!que.empty()) {
            int counts = que.size();
            for (int i = 0; i < counts; ++i) {
                auto node = que.front();
                que.pop();

                if (node->left) {
                    que.push(node->left);
                }
                if (node->right) {
                    que.push(node->right);
                }
                if (i == counts - 1) {
                    ans.push_back(node->val);
                }
            }
        }

        return ans;
    }
};
// @lc code=end
