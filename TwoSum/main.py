from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        d = {}
        for i in range(len(nums)):
            if nums[i] in d :
                d[nums[i]][i] = None
            else:
                d[nums[i]] = { i : None }
                
            if (target - nums[i]) in d:
                for y in d[target - nums[i]].keys(): 
                    if  i != y:
                        return [i,y]
        
print( Solution().twoSum([2,7,9,11],9) )