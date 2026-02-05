/*
 * @lc app=leetcode.cn id=82 lang=cpp
 *
 * [82] 删除排序链表中的重复元素 II
 */

#include <stdlib.h>
#include "define.h"

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
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        ListNode* dummy = new ListNode(0, head);
        ListNode* p = dummy;
        ListNode* cur = head;
        while (cur != nullptr) {
            while (cur->next != nullptr && cur->val == cur->next->val) {
                cur = cur->next;
            }
            if (p->next == cur) {
                p = cur;
                cur = cur->next;
                continue;
            }
            p->next = cur->next;
            cur = cur->next;
        }
        return dummy->next;
    }
};
// @lc code=end
