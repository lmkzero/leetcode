/*
 * @lc app=leetcode.cn id=470 lang=cpp
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
// The rand7() API is already defined for you.
// int rand7();
// @return a random integer in the range 1 to 7

#include <cstdlib>
class Solution {
   public:
    int rand10() {
        // ①已知 rand7()可以等概率生成1~7；
        // ②那么rand7()-1可以等概率生成0~6；
        // ③那么（rand7()-1）* 7 可以等概率生成{0, 7, 14, 21, 28, 35, 42};
        // ④那么（rand7()-1）* 7 + rand7()可以等概率生成1~49。

        // 现在我们实现了等概率生成1~49的功能，我们规定如果生成的数大于40，
        // 我们就重新生成，即只要生成在1~40之间的数，这一我们就可以生成1~40的等概率分布。

        // 再接着，我们用（1~40）%10，就可以得到0~9
        // 最后，（1~40）%10 + 1 ，就可以得到1~10了。

        int ans = 0;
        do {
            ans = ((rand7() - 1) * 7 + rand7());
        } while (ans > 40);
        return 1 + ans % 10;
    }

   private:
    int rand7() { return std::rand() % (7 - 1) + 1; }
};
// @lc code=end
