/*
 * @lc app=leetcode.cn id=23 lang=cpp
 *
 * [23] 合并K个排序链表
 */

#include <stdlib.h>
#include <queue>
#include <vector>
#include "define.h"

using namespace std;

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
    ListNode *mergeKLists(vector<ListNode *> &lists) {
        if (lists.size() == 0) {
            return nullptr;
        }
        ListNode *p = lists[0];
        for (int i = 1; i < lists.size(); i++) {
            p = mergeTwoLists(p, lists[i]);
        }
        return p;
    }

    ListNode *mergeTwoLists(ListNode *l1, ListNode *l2) {
        ListNode head(-1);
        for (ListNode *p = &head; l1 != nullptr || l2 != nullptr; p = p->next) {
            int val1 = l1 == nullptr ? INT_MAX : l1->val;
            int val2 = l2 == nullptr ? INT_MAX : l2->val;
            if (val1 <= val2) {
                p->next = l1;
                l1 = l1->next;
            } else {
                p->next = l2;
                l2 = l2->next;
            }
        }
        return head.next;
    }

    ListNode *mergeKListsWithHeap(vector<ListNode *> &lists) {
        auto cmp = [](ListNode *a, ListNode *b) { return a->val > b->val; };
        priority_queue<ListNode *, vector<ListNode *>, decltype(cmp)> pq;
        for (auto head : lists) {
            if (head) {
                pq.push(head);
            }
        }
        ListNode *dummy = new ListNode();
        ListNode *cur = dummy;
        while (!pq.empty()) {
            ListNode *node = pq.top();
            pq.pop();
            if (node->next) {
                pq.push(node->next);
            }
            cur->next = node;
            cur = cur->next;
        }
        return dummy->next;
    }
};
// @lc code=end
