class NumArray:

    def __init__(self, nums):
        self.num_array = nums
        
        

    def sumRange(self, left, right):
        return sum(self.num_array[left:right + 1])
        


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)
