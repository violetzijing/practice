#!/usr/bin/env ruby
#
# https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
#
class ListNode
  attr_accessor :val, :next
  def initialize(val)
    @val = val
    @next = nil
  end
end


def remove_nth_from_end head, n
  fast = head
  slow = head

  (0..n-1).each { fast = fast.next }

  return head.next if fast == nil

  while fast.next != nil
    fast = fast.next
    slow = slow.next
  end

  slow.next = slow.next.next

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
#n2 = ListNode.new 2
#n3 = ListNode.new 3
#n4 = ListNode.new 4
#n5 = ListNode.new 5
#n6 = ListNode.new 6
#n7 = ListNode.new 7
#
#n1.next = n2
#n2.next = n3
#n3.next = n4
#n4.next = n5
#n5.next = n6
#n6.next = n7
#
print_lst n1
puts
print_lst remove_nth_from_end n1, 1
