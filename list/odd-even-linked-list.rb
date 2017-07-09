#!/usr/bin/env ruby
#
# https://leetcode.com/problems/odd-even-linked-list/#/description
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

def odd_even_list(head)
  return if head == nil
  return head if head.next == nil or head.next.next == nil
  even_head = head.next

  odd = head
  even = even_head
  odd_pre = nil
  while even != nil
    odd.next = even.next
    odd_pre = odd
    odd = even.next
    if odd == nil
      even = nil
    else
      even.next = odd.next
      even = odd.next
    end
  end

  if odd == nil
    odd_pre.next = even_head
  else
    odd.next = even_head
  end

  return head
end

n1 = ListNode.new 1
#n2 = ListNode.new 2
#n3 = ListNode.new 3
#n4 = ListNode.new 4
#n5 = ListNode.new 5

#n1.next = n2
#n2.next = n3
#n3.next = n4
#n4.next = n5

print_lst odd_even_list n1

