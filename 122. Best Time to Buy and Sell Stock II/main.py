from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        ln = len(prices)
        if ln < 2:
            return 0
        prof = 0
        has = False
        for i in range(ln-1):
            if prices[i] < prices[i+1] and not has:
                has = True
                prof-=prices[i]
            if prices[i] > prices[i+1] and has:
                has = False
                prof+=prices[i]
        if has :
            prof+=prices[-1]
        return prof