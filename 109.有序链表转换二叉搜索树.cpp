/*
 * @lc app=leetcode.cn id=109 lang=cpp
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
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
    TreeNode* sortedListToBST(ListNode* head) {
        // 自底向上
        int len=0;
        ListNode *p=head;
        while(p){
            len++;
            p=p->next;
        }
        return sortedListToBST(head,0,len-1);
    }

private:
    TreeNode* sortedListToBST(ListNode* &list,int start,int end){
        if(start>end){
            return nullptr;
        }
        int mid=start+(end-start)/2;
        TreeNode *leftChild=sortedListToBST(list,start,mid-1);
        TreeNode *parent=new TreeNode(list->val);
        parent->left=leftChild;
        list=list->next;
        parent->right=sortedListToBST(list,mid+1,end);
        return parent;
    }
};
// @lc code=end

