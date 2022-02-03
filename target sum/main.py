from typing import List

class Solution:
    # def findTargetSumWays(self, nums: List[int], target: int) -> int:
    #     if len(nums) == 1 :
    #         if target == 0 and nums[0] == 0:
    #             return 2
    #         if target == nums[0] or target == nums[0] * -1:
    #             return 1
    #         else:
    #             return 0
    #     if nums[-1] == 0:
    #         return 2 * self.findTargetSumWays(nums[:-1],target)

    #     return self.findTargetSumWays(nums[:-1],target-nums[-1]) + self.findTargetSumWays(nums[:-1],target+nums[-1])
        # def findTargetSumWays(self, nums: List[int], target: int) -> int:
        # c=0
        # def helper(i,t):
        #     nonlocal c
        #     if i == len(nums):
        #         if t == 0:
        #             c+=1
        #     else:
        #         helper(i+1,t-nums[i])
        #         helper(i+1,t+nums[i])
        # helper(0,target)
        # return c
    # def findTargetSumWays(self, nums: List[int], target: int) -> int:
    #     cache = { x : {} for x in range(len(nums)) }
    #     def helper(i,t):
    #         if i == len(nums):
    #             if t == 0:
    #                 return 1
    #             else:
    #                 return 0
    #         else:
    #             nonlocal cache
    #             s = t-nums[i]
    #             if i in cache and s in cache[i]:
    #                 return cache[i][s]
    #             sum =  helper(i+1,s)  +helper(i+1,t+nums[i])
    #             cache[i][s] = sum
    #             return sum
    #     return helper(0,target)
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        cache = { x : {} for x in range(len(nums) ) }
        def dp( nums: List[int], target: int):
  
            i = len(nums)-1
            if target in cache[i]:
                return cache[i][target]
            
            if len(nums) == 1 :
                if target == 0 and nums[0] == 0:
                    return 2
                if target == nums[0] or target == nums[0] * -1:
                    return 1    
                return 0

            cache[i][target] =  dp(nums[:-1],target-nums[-1]) + dp(nums[:-1],target+nums[-1])
            return  cache[i][target]
        return dp(nums , target)
    
print(Solution().findTargetSumWays([22,36,7,44,38,32,16,32,1,16,25,45,49,45,27,9,41,31,10,15],1))
# print(Solution().findTargetSumWays([2,20,24,38,44,21,45,48,30,48,14,9,21,10,46,46,12,48,12,38] , 48))