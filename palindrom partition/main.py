from typing import List

class Solution:
    def partition(self, s: str) -> List[List[str]]:
        result , l_s  = [] , len(s)
        def helper(pos , c_r):
            if pos >= l_s:
                result.append(c_r[:])
                return
            for i in range(pos,l_s):
                if self.isPal(s[pos:i+1]):
                    c_r.append(s[pos:i+1])
                    helper(i+1 , c_r)
                    c_r.pop()
        helper(0,[])
        return result
            

    def isPal(self,s):
        for i in range(len(s)//2):
            if s[i] != s[-(1+i)]:
                return False
        return True

print(Solution().partition("abc"))