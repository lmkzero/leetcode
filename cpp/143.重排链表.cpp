/*
 * @lc app=leetcode.cn id=143 lang=cpp
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
   public:
    void reorderList(ListNode* head) {
        if (head == nullptr || head->next == nullptr) {
            return;
        }
        // 找到中间结点
        ListNode* slow = head;
        ListNode* fast = head;
        ListNode* prev = nullptr;
        while (fast && fast->next) {
            prev = slow;
            slow = slow->next;
            fast = fast->next->next;
        }
        // 截断链表
        prev->next = nullptr;
        // reverse
        slow = reverse(slow);
        // merge
        ListNode* cur = head;
        while (cur->next) {
            ListNode* tmp = cur->next;
            cur->next = slow;
            slow = slow->next;
            cur->next->next = tmp;
            cur = tmp;
        }
        cur->next = slow;
    }

    ListNode* reverse(ListNode* head) {
        if (head == nullptr || head->next == nullptr) return head;
        ListNode* prev = head;
        for (ListNode *curr = head->next, *next = curr->next; curr;
             prev = curr, curr = next, next = next ? next->next : nullptr) {
            curr->next = prev;
        }
        head->next = nullptr;
        return prev;
    }
};
// @lc code=end
