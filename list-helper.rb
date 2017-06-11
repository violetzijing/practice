#!/usr/bin/env ruby

class ListNode
  attr_accessor :val, :next
  def initialize(val)
    @val = val
    @next = nil
  end
end

def print_lst head
  node = head
  while node != nil
    print "#{node.val} -> "
    node = node.next
  end
end

def create_lst array
  head = ListNode.new array[0]
  parent = head
  i = 1
  while i < array.length
    node = ListNode.new array[i]
    parent.next = node
    parent = node
    i += 1
  end

  return head
end
