from typing import List

class Solution:
    # def countBits(self, n: int) -> List[int]:
    #     def helper(s):
    #         r = 0
    #         while True:
    #             if s == 0:
    #                 return r
    #             if s % 2 == 1:
    #                 r+=1
    #             s = s // 2

    #     return [helper(x) for x in range(n+1)]
    def countBits(self, n: int) -> List[int]:
        mem = [0 for _ in range(n+1)]
        def helper(s):
            if s %2 == 0:
                mem[s]=mem[s//2]
            else:
                mem[s]=mem[s//2] + 1
            return mem[s]

        return [helper(x) for x in range(n+1)]

print(Solution().countBits(20))

        

    