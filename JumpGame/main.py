from typing import List 

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) <=1:
            return True
        c= 0
        for x in range(len(nums) -2 , -1 , -1):
            c+=1
            if nums[x] >= c:
                c = 0
        return c == 0

print(Solution().canJump([2,0,0]))