from typing import List
from operator import itemgetter

class Solution:
    def reconstructQueue(self, people: List[List[int]]) -> List[List[int]]:
        people.sort(key = lambda x: (x[0],-x[1]) )
        # s1 = sorted(s1,key = itemgetter(0))
        people = [[*x,x[1]] for x in people]
        ret  = []
        while True:
            for i in range(len(people)):
                if people[i][2] == 0:
                    ret.append(people[i][0:2])
                    for y in range(i-1,-1,-1):
                        people[y][2] = people[y][2]-1
                    del people[i]
                    break
            if len(people) == 0:
                return ret
    
    #  def reconstructQueue(self, people: List[List[int]]) -> List[List[int]]:
    #     people.sort(key=lambda x:(x[0],-x[1]),reverse=True) #sorting by height first in descending order and if height is equal then by ascending order of k
    #     ans=[]
    #     for i in people:
    #         ans.insert(i[1],i) #insert at given index from max height to min height
    #     return ans


print(Solution().reconstructQueue([[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]))
print(Solution().reconstructQueue([[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]))