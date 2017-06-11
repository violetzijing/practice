#!/usr/bin/env ruby
#
# https://leetcode.com/problems/reverse-linked-list/#/description
#

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

def reverse_between head, m, n
  return head if m == n
  return if head == nil

  p1 = nil
  p2 = nil
  p3 = nil
  p4 = nil

  parent = nil
  node = head
  i = 1
  while node != nil and i < m
    parent = node
    node = node.next
    i += 1
  end
  p1 = parent
  p3 = node

  while node != nil and i <= n
    tmp = node.next
    node.next = parent
    parent = node
    node = tmp
    i += 1
  end

  p4 = parent
  p2 = node

  if p1 == nil
    head = p4
  else
    p1.next = p4
  end
  p3.next = p2

  return head
end

n1 = ListNode.new 1
n2 = ListNode.new 2
n3 = ListNode.new 3
n4 = ListNode.new 4
n5 = ListNode.new 5

n1.next = n2
n2.next = n3
n3.next = n4
n4.next = n5

print_lst reverse_between n5, 1, 2
