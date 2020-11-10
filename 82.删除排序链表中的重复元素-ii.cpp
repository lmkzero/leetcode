/*
 * @lc app=leetcode.cn id=82 lang=cpp
 *
 * [82] 删除排序链表中的重复元素 II
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
        ListNode* p=head->next;
        if(head->val==p->val){
            while(p&&head->val==p->val){
                ListNode* tmp=p;
                p=p->next;
                delete tmp;
            }
            delete head;
            return deleteDuplicates(p);
        }else{
            head->next=deleteDuplicates(head->next);
            return head;
        }
    }
};
// @lc code=end

