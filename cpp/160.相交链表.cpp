/*
 * @lc app=leetcode.cn id=160 lang=cpp
 *
 * [160] 相交链表
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
    int getLength(ListNode *head) {
        int num = 0;
        for (; head; head = head->next) {
            num++;
        }
        return num;
    }

    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        if (headA == nullptr || headB == nullptr) {
            return nullptr;
        }
        int lenA = getLength(headA);
        int lenB = getLength(headB);
        int len = 0;
        if (lenA > lenB) {
            len = lenB;
            while (lenA > lenB) {
                headA = headA->next;
                lenA--;
            }
        } else {
            len = lenA;
            while (lenB > lenA) {
                headB = headB->next;
                lenB--;
            }
        }
        while (len--) {
            if (headA == headB) {
                return headA;
            }
            headA = headA->next;
            headB = headB->next;
        }
        return nullptr;
    }
};
// @lc code=end
