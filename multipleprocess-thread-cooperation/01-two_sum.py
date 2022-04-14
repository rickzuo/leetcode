# -*- coding: utf-8 -*-
# @Time    : 2019/10/21 8:44 下午
# @Author  : Rick
# @File    : 01-two_sum.py
# @Software: PyCharm

class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for i in range(0,len(nums)):
            for j in range(i+1,len(nums)):
                if nums[i]+nums[j] == target:
                    return [i,j]

if __name__ == '__main__':
    """
    Given nums = [2, 7, 11, 15], target = 9,
        
        Because nums[0] + nums[1] = 2 + 7 = 9,
        return [0, 1].
    """
    nums = [2, 7, 11, 15]
    target = 9
    print(Solution().twoSum(nums,target))