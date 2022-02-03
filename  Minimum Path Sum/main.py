from typing import List

class Solution:
    # def minPathSum(self, grid: List[List[int]]) -> int:
    #     min_cost = 0
    #     def helper(cur_cost , pos_m , pos_n):
    #         nonlocal min_cost
    #         cur_cost = cur_cost + grid[pos_m][pos_n]
    #         if pos_n == len(grid[0])-1 and pos_m == len(grid)-1:
    #             if min_cost == 0:
    #                 min_cost = cur_cost
    #             elif cur_cost < min_cost:
    #                 min_cost = cur_cost
            
    #         if pos_m < len(grid)-1:
    #             helper(cur_cost , pos_m+1 , pos_n)
    #         if pos_n < len(grid[0])-1:
    #             helper(cur_cost , pos_m , pos_n+1)
    #     helper(0,0,0)
    #     return min_cost
    def minPathSum(self, grid: List[List[int]]) -> int:
        min_grid = [ [None for _ in range(len(grid[0]))] for _ in range(len(grid))]
        def helper(cur_cost , pos_m , pos_n):
            cur_cost = cur_cost + grid[pos_m][pos_n]
            if pos_n == len(grid[0])-1 and pos_m == len(grid)-1:
                return cur_cost

            if min_grid[pos_m][pos_n]:
                return min_grid[pos_m][pos_n] + cur_cost

            r1 , r2 = float('inf') , float('inf')
            if pos_m < len(grid)-1:
                r1 = helper(cur_cost , pos_m+1 , pos_n)
            if pos_n < len(grid[0])-1:
                r2 = helper(cur_cost , pos_m , pos_n+1)
            r_v = min(r1,r2)
            min_grid[pos_m][pos_n] = r_v - cur_cost
            return r_v
        return helper(0,0,0)
        #     def minPathSum(self, grid: List[List[int]]) -> int:
        # m=len(grid)-1
        # n=len(grid[0])-1
        
        # def dp(m,n,memo={}):
        #     key=str(m)+','+str(n)
        #     if key in memo: return memo[key]
        #     if m==0 and n==0: return grid[m][n]
        #     if m<0 or n<0: return float(inf)
            
        #     memo[key]=grid[m][n]+min(dp(m-1,n,memo),dp(m,n-1,memo))
        #     return memo[key]
        # return dp(m,n)

        

print(Solution().minPathSum([[1,3,1],[1,5,1],[4,2,1]]))
print(Solution().minPathSum([[1,2,3],[4,5,6]]))
print(Solution().minPathSum([[7,1,3,5,8,9,9,2,1,9,0,8,3,1,6,6,9,5],[9,5,9,4,0,4,8,8,9,5,7,3,6,6,6,9,1,6],[8,2,9,1,3,1,9,7,2,5,3,1,2,4,8,2,8,8],[6,7,9,8,4,8,3,0,4,0,9,6,6,0,0,5,1,4],[7,1,3,1,8,8,3,1,2,1,5,0,2,1,9,1,1,4],[9,5,4,3,5,6,1,3,6,4,9,7,0,8,0,3,9,9],[1,4,2,5,8,7,7,0,0,7,1,2,1,2,7,7,7,4],[3,9,7,9,5,8,9,5,6,9,8,8,0,1,4,2,8,2],[1,5,2,2,2,5,6,3,9,3,1,7,9,6,8,6,8,3],[5,7,8,3,8,8,3,9,9,8,1,9,2,5,4,7,7,7],[2,3,2,4,8,5,1,7,2,9,5,2,4,2,9,2,8,7],[0,1,6,1,1,0,0,6,5,4,3,4,3,7,9,6,1,9]]))