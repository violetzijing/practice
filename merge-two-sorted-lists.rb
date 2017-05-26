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
  if l1 == nil and l2 != nil
    return l2
  elsif l2 == nil and l1 != nil
    return l1
  elsif l1 == nil and l2 == nil
    return
  end

  top = (l1.val < l2.val ? l1 : l2)
  iterator = top

  while l1 != nil and l2 != nil

    if l1.val < l2.val
      if l1.next == nil
        l1.next = l2
        break
      end

      if l1.next.val < l2.val
        l1 = l1.next
      else
        l2_next = l2.next
        l1_next = l1.next
        l1.next = l2
        l2.next = l1_next
        if l1.next != nil
          l1 = l1_next if l2.val > l1.next.val
        end
        l2 = l2_next
      end
    else
      if l2.next == nil
        l2.next = l1
        break
      end

      if l2.next.val < l1.val
        l2 = l2.next
      else
        l1_next = l1.next
        l2_next = l2.next
        l2.next = l1
        l1.next = l2_next
        if l2.next != nil
          l2 = l2_next if l1.val > l2.next.val
        end
        l1 = l1_next
      end
    end
  end

  return top
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

nodetmp1 = ListNode.new(1)
nodetmp2 = ListNode.new(1)
print_list(merge_two_lists(nodetmp1, nodetmp2))
