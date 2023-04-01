/*
 * @lc app=leetcode.cn id=138 lang=cpp
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* next;
    Node* random;
    
    Node(int _val) {
        val = _val;
        next = NULL;
        random = NULL;
    }
};
*/

class Solution {
public:
    Node* copyRandomList(Node* head) {
        // copy的新结点添加到原节点的后面
        // 拆分链表
        if(head==nullptr){
            return head;
        }
        for (Node* cur = head; cur != nullptr; ) {
            Node* node = new Node(cur->val);
            node->next = cur->next; 
            cur->next = node; 
            cur = node->next; 
        } 
        for (Node* cur = head; cur != nullptr; ) { 
            if (cur->random != NULL) 
                cur->next->random = cur->random->next; 
            cur = cur->next->next; 
        } 
        Node dummy(-1); 
        for (Node* cur = head, *new_cur = &dummy; cur != nullptr; ) { 
            new_cur->next = cur->next; 
            new_cur = new_cur->next; 
            cur->next = cur->next->next; 
            cur = cur->next; 
        } 
        return dummy.next;
    }
};
// @lc code=end

