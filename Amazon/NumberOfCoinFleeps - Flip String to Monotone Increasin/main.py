# def minCoins(data):
#     def helper(data, total_sum, val):
#         if len(val) == 0:
#             return total_sum

#         cur_it = val[0]
#         cur_sym = cur_it[0]
#         cur_sum = cur_it[1]

#         reduced = val[1:]

#         if val != cur_sym:
#             total_sum += cur_sym
#             if val == 'H':
#                 cur_sym = 'T'
#             else:
#                 cur_sym = 'H'

#         sum_1 = helper(reduced, total_sum, cur_sum, 'H')
#         sum_2 = helper(reduced, total_sum, cur_sum, 'T')

#         if sum_1 < sum_2:
#             return sum_1
#         return sum_2
    
#     if len(data) == 0:
#         return 0

#     calculated = []
#     prev_val = 'H'
#     prev_idx = 0
#     for idx, val in enumerate(list(data)):
#         if val != prev_val:
#             calculated.append([prev_val, idx - prev_idx])
#             prev_val = val
#             prev_idx = idx
#     if calculated[0][1] == 0:
#         calculated = calculated[1:]

#     start = calculated[0][0]

#     sum_1 = helper(calculated, 0, start, 'H')
#     sum_2 = helper(calculated, 0, start, 'T')
#     if sum_1 < sum_2:
#         return sum_1
#     return sum_2




# minCoins('THHTTHTTHT')

# # minimumCoinFlips("THHHTH")


class Solution:
    # def helper(self,s:str):
    #     c_0 , c_1 = 0 , 0
    #     m_1 = False
    #     wr = False
    #     for p , i in enumerate(s):
    #         if not m_1 and i == '1':
    #             m_1 = True
    #         if m_1 and i == '0':
    #             wr = True
    #         if wr and i == '0':
    #             r_0 , r_1 = self.helper(s[p:])
    #             return (c_0 + r_0,c_1 + r_1)
    #         if m_1:
    #             c_1+=1
    #         if wr:
    #             c_0+=1
    #     return (c_0,c_1)      
    # def minFlipsMonoIncr(self, s: str) -> int:
    #     r1 , r2 =  self.helper(s)
    #     return min(r1,r2)
    def minFlipsMonoIncr(self, s: str) -> int:
        minCostEnd1 , minCostEnd0 = 0 , 0
        for i in s:
            if i == '0':
                minCostEnd1=min(minCostEnd0,minCostEnd1) +1 
            else:
                minCostEnd1 = min(minCostEnd0,minCostEnd1)
                minCostEnd0+=1
    

        return min(minCostEnd1 , minCostEnd0)
s = "00011000"
a = Solution().minFlipsMonoIncr(s)
print(a)