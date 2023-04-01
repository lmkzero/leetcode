/*
 * @lc app=leetcode.cn id=203 lang=cpp
 *
 * [203] 移除链表元素
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
class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        if(head==nullptr){
            return head;
        }
        ListNode* dummy=new ListNode(val+1);
        dummy->next=head;
        ListNode* pre=dummy;
        ListNode* cur=head;
        while(cur){
            if(cur->val==val){
                ListNode* tmp=cur;
                pre->next=cur->next;
                cur=pre->next;
                delete tmp;
            }else{
                pre=pre->next;
                cur=pre->next;
            }
        }
        return dummy->next;
    }
};
// @lc code=end

