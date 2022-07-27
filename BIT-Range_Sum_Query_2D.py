# let's implement Range Sum Query on 2D matrix (LEETCODE Locked)
# Good problems and solutions can be found here - https://evanyang.gitbooks.io/leetcode/content/SUMMARY.html
class BIT:
    def __init__(self,matrix):
        self.m,self.n = len(matrix),len(matrix[0])
        self.tree = [[0]*(self.n+1) for _ in range(self.m+1)]
        self.nums = [[0]*(self.n) for _ in range(self.m)]
        self.matrix = matrix
        
        self.construct()
    
    def construct(self):
        r,c = 0,0
        while r<self.m:
            c = 0
            while c<self.n:
                self.update(r,c,self.matrix[r][c])
                c += 1   
            r += 1
                
        # print(*self.tree,sep="\n")
    
    def update(self,r,c,newVal):
        delta = newVal - self.nums[r][c]
        self.nums[r][c] = newVal
        
        i,j = r+1,c+1
        while i<=self.m:
            j = c+1
            while j<=self.n:
                self.tree[i][j] += delta
                j += j & (-j)
                
            i += i & (-i)
     
    #   .  .  .   . . .  .            
    #   .  +  .   . . .  -
    #   . (r1,c1) . . .  .
    #   .  .  .   . . .  .
    #   .  .  .   . . .  .
    #   .  -  .   . . . (r2,c2)
    
    def sumRegion(self,r1,c1,r2,c2):
        return  self.getSum(r2+1,c2+1) - self.getSum(r1,c2+1) - self.getSum(r2+1,c1) + self.getSum(r1,c1)
        
    def getSum(self,r,c):
        sum = 0
        i,j = r,c
        while i>0:
            j = c
            while j>0:
                sum += self.tree[i][j]
                j -= j & (-j)
                
            i -= i & (-i)
        return sum
        
matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

bit = BIT(matrix)
print(bit.sumRegion(2, 1, 4, 3))
bit.update(3, 2, 2)
print(bit.sumRegion(2, 1, 4, 3)) 
print(bit.sumRegion(1, 1, 2, 2))

# Output:
# 8
# 10
# 11
# Intuition from - https://evanyang.gitbooks.io/leetcode/content/LeetCode/range_sum_query_2d_-_mutable.html
