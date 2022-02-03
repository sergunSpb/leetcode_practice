from typing import List

class SolutionBrutForce:
    def subarraySum(self, nums: List[int], k: int) -> int:
        sums = { 0 : 0}
        sum = 0
        for i in range(0,len(nums)):
            sum+=nums[i]
            sums[i+1] = sum
        answer = 0
        for st in range(0,len(nums)):
            for en in range(st+1,len(nums)+1):
                if sums[en] - sums[st] == k:
                    answer+=1
        return answer

# print(SolutionBrutForce().subarraySum([1,2,3],3))      

class SolutionBrutForce2:
    def subarraySum(self, nums: List[int], k: int) -> int:
        answer = 0
        for i in range(0,len(nums)):
            sum = 0
            for y in range(i,len(nums)):
                sum += nums[y]
                if sum == k:
                    answer+=1
        return answer

# print(SolutionBrutForce2().subarraySum([1,2,3],3))      

class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        answer = 0
        s = { 0:1 }
        sum = 0
        for i in range(0,len(nums)):
            sum+=nums[i]
            if sum - k in s :
                answer+=1
                s[sum-k] +=1
            if sum in s:
                 s[sum]+=1
            else:
                s[sum] = 1
        return answer

print(Solution().subarraySum([1,2,3],3))      
        