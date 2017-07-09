#!/usr/bin/env ruby
#
# https://leetcode.com/problems/same-tree/#/description
#

class TreeNode
  attr_accessor :val, :left, :right
  def initialize(val)
    @val = val
    @left, @right = nil, nil
  end
end

def is_same_tree p, q
  return true if p == nil and q == nil
  return false if p == nil and q != nil
  return false if p != nil and q == nil

  in_order_traverse_without_recursion p, q
end

def in_order_traverse_without_recursion p, q
  stack1 = []
  stack2 = []
  stack1.push p if p != nil
  stack2.push q if q != nil

  while stack1.length > 0 and stack2.length > 0
    tmp_p = stack1.last
    tmp_q = stack2.last

    stack1.pop
    stack2.pop

    if tmp_p != nil and tmp_q != nil
      return false if tmp_p.val != tmp_q.val
      stack1.push tmp_p.left
      stack1.push tmp_p.right
      stack2.push tmp_q.left
      stack2.push tmp_q.right
    end

    return false if tmp_p != nil and tmp_q == nil
    return false if tmp_p == nil and tmp_q != nil
    #return false if stack1.length != stack2.length
  end

  return stack1.length == stack2.length
end

node1 = TreeNode.new 1
node2 = TreeNode.new 2
node3 = TreeNode.new 3

node4 = TreeNode.new 2
node5 = TreeNode.new 3
node6 = TreeNode.new 4

node1.right = node3

node4.left = node5
node4.right = node6

puts is_same_tree node1, node4
