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

l1 = int2lst 5
l2 = int2lst 5

def chop_lst lst
  while lst.next.next.nil? and lst.next.val == 0 do
    lst.next = nil
    lst = lst.next
  end
end

def add_two_numbers(l1, l2)
  l3 = ListNode.new(0)
  start = l3
  loop do
    l3_next = ListNode.new(0)
    l3.next = l3_next

    l1.nil? ? val1 = 0 : val1 = l1.val
    l2.nil? ? val2 = 0 : val2 = l2.val

    l3.val += val1 + val2
    if l3.val >= 10
      l3.val -= 10
      l3_next.val += 1
    end

    l1 = l1.next unless l1.nil?
    l2 = l2.next unless l2.nil?

    if l1.nil? and l2.nil?
      l3.next = nil if l3_next.val == 0
      break
    else
      l3 = l3_next
    end
  end

  return start
end

lst = add_two_numbers l1, l2
print_lst(lst)
