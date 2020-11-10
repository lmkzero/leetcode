/*
 * @lc app=leetcode.cn id=19 lang=cpp
 *
 * [19] 删除链表的倒数第N个节点
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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode dummy(-1);
        dummy.next=head;
        ListNode* p=&dummy;
        ListNode* q=&dummy;
        for(int i=0;i<n;i++){
            q=q->next;
        }
        while(q->next){
            p=p->next;
            q=q->next;
        }
        ListNode* tmp=p->next;
        p->next=p->next->next;
        delete tmp;
        return dummy.next;
    }
};
// @lc code=end

