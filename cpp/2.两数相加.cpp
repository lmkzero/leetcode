/*
 * @lc app=leetcode id=2 lang=cpp
 *
 * [2] Add Two Numbers
 */

#include <iostream>
#include <vector>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        if (l1 == NULL && l2 == NULL) {
            return NULL;
        } else if (l1 == NULL) {
            return l2;
        } else if (l2 == NULL) {
            return l1;
        }

        int a = l1->val + l2->val;
        ListNode* p = new ListNode(a % 10);
        p->next = addTwoNumbers(l1->next, l2->next);
        if (a >= 10) {
            p->next = addTwoNumbers(p->next, new ListNode(1));
        }
        return p;
        // ListNode* head=NULL;
        // ListNode* index=NULL;
        // ListNode* index1=l1;
        // ListNode* index2=l2;
        // int add=0;
        // int res=0;
        // while(true){
        //     if(index1==NULL&&add==0){
        //         return head;
        //     }else if(add!=0){
        //         ListNode* temp=new ListNode(res);
        //         index->next=temp;
        //         index=index->next;
        //     }else{
        //         int sum=index1->val+index2->val+add;
        //         if(sum >= 10){
        //             res=sum-10;
        //             add=1;
        //         }else{
        //             res=sum;
        //             add=0;
        //         }
        //         if(head==NULL){
        //             head=new ListNode(res);
        //             index=head;
        //         }else{
        //             ListNode* temp=new ListNode(res);
        //             index->next=temp;
        //             index=index->next;
        //         }
        //         index1=index1->next;
        //         index2=index2->next;
        //     }
        // }
    }
};
// @lc code=end
