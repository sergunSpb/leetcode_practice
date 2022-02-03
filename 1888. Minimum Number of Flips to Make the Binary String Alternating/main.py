class Solution:
    def helper(self,s) -> int:
        if s in self.cache :
            return self.cache[s]
        
        minT2 = 0
        pr = None
        for i in s:
            if pr and pr == i:
                if i == '0':
                    pr = '1'
                else:
                    pr = '0'
                minT2 +=1
            else:
                pr =i
        self.cache[s] = minT2
        return minT2
        
    def minFlips(self, s: str) -> int:
        if s == 1:
            return 0
        self.cache = {}

        l = len(s)
        r = s
        m = None
        for i in range(0,l-1):
            s= self.helper(r)
            if m is None or m > s:
                m = s
            r = r[1:] + r[0]
        return m
                