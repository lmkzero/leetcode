/*
 * @lc app=leetcode.cn id=108 lang=cpp
 *
 * [108] 将有序数组转换为二叉搜索树
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
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        return sortedArrayToBST(nums.begin(),nums.end());
    }

    template<typename RandomAccessIterator>
    TreeNode* sortedArrayToBST(RandomAccessIterator first,RandomAccessIterator last){
        auto length=distance(first,last);
        if(length<=0){
            return nullptr;
        }
        auto mid=first+length/2;
        TreeNode* root=new TreeNode(*mid);
        root->left=sortedArrayToBST(first,mid);
        root->right=sortedArrayToBST(mid+1,last);
        return root;
    }
};
// @lc code=end

