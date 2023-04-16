/*
 * @lc app=leetcode.cn id=155 lang=cpp
 *
 * [155] 最小栈
 */

// @lc code=start
class MinStack {
   private:
    stack<int> numStack;
    stack<int> minStack;

   public:
    MinStack() {}

    void push(int val) {
        if (numStack.size() == 0) {
            minStack.push(val);
        }
        if (val <= minStack.top()) {
            minStack.push(val);
        }
        numStack.push(val);
    }

    void pop() {
        if (numStack.top() == minStack.top()) {
            minStack.pop();
        }
        numStack.pop();
    }

    int top() { return numStack.top(); }

    int getMin() { return minStack.top(); }
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(val);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */
// @lc code=end
