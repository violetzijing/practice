#!/usr/bin/env ruby
# https://leetcode.com/problems/add-two-numbers/#/description
#
class ListNode
  attr_accessor :val, :next
  def initialize(val)
      @val = val
      @next = nil
  end
end

def int2lst(num)
  head = ListNode.new(nil)
  node = head
  while num != 0 do
    digit = num % 10
    num /= 10
    node.next = ListNode.new(digit)
    node = node.next
  end
  return head.next
end

def print_lst(lst)
  while lst != nil do
    print "#{lst.val} -> "
    lst = lst.next
  end
  puts ""
end

l1 = int2lst 6
l2 = int2lst 789

def chop_lst lst
  while lst.next.next.nil? and lst.next.val == 0 do
    lst.next = nil
    lst = lst.next
  end
end

def add_two_numbers(l1, l2)
  head = ListNode.new nil
  prev = head

  carry = 0
  until l1.nil? and l2.nil? do
    val1 = l1.nil? ? 0 : l1.val
    val2 = l2.nil? ? 0 : l2.val

    sum = val1 + val2 + carry
    carry = 0

    if sum >= 10
      sum -= 10
      carry = 1
    end

    prev.next = ListNode.new sum

    l1 = l1.next unless l1.nil?
    l2 = l2.next unless l2.nil?
    prev = prev.next
  end
  
  if carry == 1
    prev.next = ListNode.new 1
  end

  return head.next
end

lst = add_two_numbers l1, l2
print_lst(lst)
