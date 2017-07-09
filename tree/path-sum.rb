#!/usr/bin/env ruby
#
# https://leetcode.com/problems/path-sum/#/description
#

class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end

def has_path_sum(root, sum)
  return false if root == nil

  return sum == root.val if root.left == nil and root.right == nil

  return has_path_sum(root.left, sum - root.val) || has_path_sum(root.right, sum - root.val)
end

node1 = TreeNode.new 1
node2 = TreeNode.new 2
node3 = TreeNode.new 3

node4 = TreeNode.new 2
node5 = TreeNode.new 3
node6 = TreeNode.new 4

node1.right = node3
node3.right = node4

node4.left = node5
node4.right = node6

puts has_path_sum node4, 20
