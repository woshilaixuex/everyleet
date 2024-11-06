#include <iostream>
#include <vector>
using namespace std;
/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: https://leetcode.cn/problems/minimum-size-subarray-sum/solutions/305704/chang-du-zui-xiao-de-zi-shu-zu-by-leetcode-solutio/?envType=study-plan-v2&envId=top-interview-150 209. 长度最小的子数组
 * 输入：target = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 没错我开始写cpp了
 * @Date: 2024-09-26 17:20
 */
class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        if (nums.size() == 0) return 0;
        int ans = INT_MAX;
        int start = 0,end = 0;
        int sum = 0;
        while (end <= nums.size()-1) {
            sum += nums[end];
            if (sum >= target){
                ans = min(ans,end - start);
                while(start <= end){
                    sum -= nums[start];
                    if (sum < target){
                        sum += nums[start];
                        break;
                    }else{
                        start++;
                        ans = min(ans,end - start);
                    }
                }
            }
            end++;
        }
        return ans == INT_MAX? 0:ans;
    }
};
