/*
 * @lc app=leetcode.cn id=142 lang=cpp
 *
 * [142] 环形链表 II
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

#include "define.h"

class Solution {
   public:
    ListNode* detectCycle(ListNode* head) {
        ListNode* slow1 = head;
        ListNode* fast = head;
        while (fast && fast->next) {
            slow1 = slow1->next;
            fast = fast->next->next;
            if (slow1 == fast) {
                ListNode* slow2 = head;
                while (slow1 != slow2) {
                    slow1 = slow1->next;
                    slow2 = slow2->next;
                }
                return slow2;
            }
        }
        return nullptr;
    }
};
// @lc code=end
