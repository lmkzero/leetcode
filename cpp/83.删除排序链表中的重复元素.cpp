/*
 * @lc app=leetcode.cn id=83 lang=cpp
 *
 * [83] 删除排序链表中的重复元素
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
    ListNode* deleteDuplicates(ListNode* head) {
        if(head==NULL||head->next==NULL){
            return head;
        }
        ListNode dummy(head->val + 1); 
        dummy.next = head; 
        recur(&dummy, head); 
        return dummy.next;

    }
    void recur(ListNode *prev, ListNode *cur) { 
        if (cur == nullptr) 
            return; 
        if (prev->val == cur->val) {
            prev->next = cur->next; 
            delete cur; 
            recur(prev, prev->next); 
        } else { 
            recur(prev->next, cur->next); 
        } 
    }
};
// @lc code=end

