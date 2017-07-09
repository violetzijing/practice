#!/usr/bin/env ruby
#
# https://leetcode.com/problems/merge-two-sorted-lists/#/description
#

require 'pry'

class ListNode
  attr_accessor :val, :next
  def initialize(val)
    @val = val
    @next = nil
  end
end

def merge_two_lists(l1, l2)
  if l1.nil?
    return l2
  elsif l2.nil?
    return l1
  end

  head = ListNode.new 0

  node = head

  while l1 != nil and l2 != nil
    if l1.val < l2.val
      node.next = l1
      l1 = l1.next
    else
      node.next = l2
      l2 = l2.next
    end
    node = node.next
  end
  
  remain = (l1 == nil) ? l2 : l1

  node.next = remain

  return head.next
end

def print_list list
  while list != nil
    print "#{list.val}, "
    list = list.next
  end
end

node1 = ListNode.new(1)
node2 = ListNode.new(4)
node3 = ListNode.new(9)

node4 = ListNode.new(2)
node5 = ListNode.new(3)
node6 = ListNode.new(5)
node7 = ListNode.new(6)

node1.next = node2
node2.next = node3

node4.next = node5
node5.next = node6
node6.next = node7

nodetmp1 = ListNode.new(2)
nodetmp2 = ListNode.new(1)
print_list(merge_two_lists(node1, node4))
