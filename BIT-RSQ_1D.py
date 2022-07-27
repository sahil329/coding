# Good resource about understanding - https://www.youtube.com/watch?v=uSFzHCZ4E-8
class BIT:
    def __init__(self,nums):
        self.n = len(nums)
        self.tree = ["*"]+copy.deepcopy(nums)
        self.construct()
        
    def construct(self):
        for i in range(1,self.n+1):
            p = i + (i & (-i))
            if p<=self.n:
                self.tree[p] += self.tree[i]
        # print(self.tree)
    
    def update(self,i,delta):
        while i<=self.n:
            self.tree[i] += delta
            i += (i & (-i))
        # print(self.tree,delta)
        
    def getSum(self,i):
        sum = 0
        while i>0:
            sum += self.tree[i]
            i -= (i & (-i))   
        return sum
