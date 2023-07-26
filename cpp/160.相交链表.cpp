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

#include "define.h"

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

    ListNode *getIntersectionNodeByCombine(ListNode *headA, ListNode *headB) {
        // p1 指向 A 链表头结点，p2 指向 B 链表头结点
        ListNode *p1 = headA, *p2 = headB;
        while (p1 != p2) {
            // p1 走一步，如果走到 A 链表末尾，转到 B 链表
            if (p1 == nullptr) {
                p1 = headB;
            } else {
                p1 = p1->next;
            }
            // p2 走一步，如果走到 B 链表末尾，转到 A 链表
            if (p2 == nullptr) {
                p2 = headA;
            } else {
                p2 = p2->next;
            }
        }
        return p1;  // 返回交点
    }
};
// @lc code=end
