#!/usr/bin/env ruby
#
# https://leetcode.com/problems/balanced-binary-tree/#/description
#

class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end

def is_balanced(root)
  return true if root == nil
  return depth(root) != -1
end

def depth node
  return 0 if node == nil

  left_height = depth node.left
  return -1 if left_height == -1

  right_height = depth node.right
  return -1 if right_height == -1

  return -1 if (left_height - right_height).abs > 1

  return [left_height, right_height].max + 1
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

puts is_balanced node4
