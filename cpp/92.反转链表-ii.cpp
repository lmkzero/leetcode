/*
 * @lc app=leetcode.cn id=92 lang=cpp
 *
 * [92] 反转链表 II
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
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        // [3,5]
        // 1
        // 2
        // if (head == NULL || head->next == NULL) {
        //     return head;
        // }
        // ListNode* pre = head;
        // for (int i = 1; i<m-1; i++) {
        //     pre = pre->next;
        // }
        // ListNode* cur = pre->next;
        // for (int j = m + 1; j <= n; j++) {
        //     ListNode* n = cur->next;
        //     ListNode* nn = n->next;
        //     ListNode* pn = pre->next;
        //     //exchange
        //     pre->next = n;
        //     n->next = pn;
        //     cur->next = nn;
        // }
        // return head;
        ListNode *dummy = new ListNode(-1);
        dummy->next = head;
        ListNode *cur = dummy;
        ListNode *pre, *front, *last;
        for (int i = 1; i <= m - 1; ++i) 
            cur = cur->next;
            pre = cur;
            last = cur->next;
            for (int i = m; i <= n; ++i) {
            cur = pre->next;
            pre->next = cur->next;
            cur->next = front;
            front = cur;
        }
        cur = pre->next;
        pre->next = front;
        last->next = cur;
        return dummy->next;
    }
};
// @lc code=end

