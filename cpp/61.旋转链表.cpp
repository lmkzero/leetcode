/*
 * @lc app=leetcode.cn id=61 lang=cpp
 *
 * [61] 旋转链表
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
    ListNode* rotateRight(ListNode* head, int k) {
        if (head == nullptr || k == 0) {
            return head;
        }
        // 遍历，获取len
        int len = 1;
        ListNode* p = head;
        while (p->next) {
            len++;
            p = p->next;
        }
        k = len - k % len;
        // 先连接成环，再断开
        p->next = head;
        for (int i = 0; i < k; i++) {
            p = p->next;
        }
        head = p->next;
        p->next = nullptr;
        return head;
    }
};
// @lc code=end
