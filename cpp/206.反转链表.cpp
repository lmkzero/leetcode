/*
 * @lc app=leetcode.cn id=206 lang=cpp
 *
 * [206] 反转链表
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

#include <stdlib.h>
#include "define.h"

class Solution {
   public:
    ListNode* reverseList(ListNode* head) {
        if (head == NULL || head->next == NULL) {
            return head;
        }
        ListNode* tail = NULL;
        ListNode* p = head;
        ListNode* q = p->next;
        while (q != NULL) {
            ListNode* old = q->next;
            p->next = tail;
            q->next = p;
            tail = p;
            p = q;
            q = old;
        }
        return p;
    }
};
// @lc code=end
