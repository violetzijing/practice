#!/usr/bin/env ruby
#
# https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/#/description
#

class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end

def sorted_array_to_bst(nums)
  return if nums.length == 0
  return TreeNode.new nums.first if nums.length == 1
  length = nums.length

  middle = length / 2
  root = TreeNode.new nums[middle]
  root.left = sorted_array_to_bst(nums[0...middle])
  root.right = sorted_array_to_bst(nums[middle + 1...length])
  return root
end


a = [1, 2, 3, 5]
puts sorted_array_to_bst a
