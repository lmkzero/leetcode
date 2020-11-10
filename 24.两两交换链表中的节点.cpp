/*
 * @lc app=leetcode.cn id=24 lang=cpp
 *
 * [24] 两两交换链表中的节点
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
    ListNode* swapPairs(ListNode* head) {
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode dummy(-1);
        dummy.next=head;
        ListNode* pre=&dummy;
        ListNode* cur=pre->next;
        ListNode* next=cur->next;
        for(;next;pre=cur,cur=cur->next,next = cur ? cur->next: nullptr){
            pre->next=next;
            cur->next=next->next;
            next->next=cur;
        }
        return dummy.next;
    }
};
// @lc code=end

