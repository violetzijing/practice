#!/usr/bin/env ruby
#
# https://leetcode.com/problems/maximum-depth-of-binary-tree/#/description
#


class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end

def max_depth(root)
  return 0 if root == nil
  [max_depth(root.left), max_depth(root.right)].max + 1
end

node1 = TreeNode.new 1
node2 = TreeNode.new 2
node3 = TreeNode.new 3

node4 = TreeNode.new 2
node5 = TreeNode.new 3
node6 = TreeNode.new 4

node1.left = node2
node1.right = node3

node4.left = node5
node5.right = node6

puts max_depth node4
