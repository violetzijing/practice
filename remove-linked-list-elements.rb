#!/usr/bin/env ruby
#
# https://leetcode.com/problems/remove-linked-list-elements/#/description
#
class ListNode
    attr_accessor :val, :next
    def initialize(val)
        @val = val
        @next = nil
    end
end
def remove_elements(head, val)
  return if head == nil

  print_lst head
  node = head
  parent = []
  while node.next != nil
    if node.next.val == val
      parent << node
    end
    node = node.next
  end

  return head if parent.length == 0

  parent.each do |p|
    p.next = p.next.next
  end

  head = head.next if head.val == val

  return head
end

def print_lst head
  node = head
  while node != nil
    print "#{node.val} -> "
    node = node.next
  end
end

n1 = ListNode.new 1
#n7 = ListNode.new 6
n2 = ListNode.new 1
#n3 = ListNode.new 3
#n4 = ListNode.new 4
#n5 = ListNode.new 5
#n6 = ListNode.new 6

n1.next = n2

#n2.next = n3
#n3.next = n4
#n4.next = n5
#n5.next = n6
#n6.next = n7

print_lst remove_elements n1, 1
