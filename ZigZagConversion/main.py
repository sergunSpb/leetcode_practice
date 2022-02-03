class Solution:
    # def convert(self, s, numRows):
    #     res = ""
    #     if numRows ==1:
    #         return s
    #     x = len(s)//numRows
    #     st  = (2*numRows -2)
    #     for i in range(numRows):
    #         for y in range(x+1):     
    #             z1 = i + st*y
    #             if z1 < len(s):
    #                 res+=s[z1]
    #             z = (y+1)*st - i
    #             if z % st != 0 and  z != z1 and z < len(s):
    #                 res+=s[z]
    #     return res
      def convert(self, s, numRows):
        if numRows ==1:
            return s
        rows = ["" for _ in range(numRows)]
        dir ,c ,p  = 1 ,0 ,0
    
        while c < len(s):
            rows[p] +=s[c]
            if (p == numRows -1 and dir == 1) or (p == 0 and dir ==-1):
                dir *=-1
            p+=dir
            c+=1
        return "".join(rows)    
print(Solution().convert("PAYPALISHIRING",3))
print(Solution().convert("P",3))