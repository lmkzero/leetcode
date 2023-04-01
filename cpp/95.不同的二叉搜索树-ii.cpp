/*
 * @lc app=leetcode.cn id=95 lang=cpp
 *
 * [95] 不同的二叉搜索树 II
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
    vector<TreeNode*> generateTrees(int n) {
        if(n==0){
            vector<TreeNode*> subTree;
            return subTree;
        }
        return generate(1,n);
    }

private:
    vector<TreeNode*> generate(int start,int end){
        vector<TreeNode*> subTree;
        if(start>end){
            subTree.push_back(nullptr);
            return subTree;
        }
        for(int k=start;k<=end;k++){
            vector<TreeNode*> leftSubs=generate(start,k-1);
            vector<TreeNode*> rightSubs=generate(k+1,end);
            for(auto i:leftSubs){
                for(auto j:rightSubs){
                    TreeNode *node=new TreeNode(k);
                    node->left=i;
                    node->right=j;
                    subTree.push_back(node);
                }
            }
        }
        return subTree;
    }
};
// @lc code=end

