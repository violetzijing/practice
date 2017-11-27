#!/usr/bin/env ruby
#
# https://leetcode.com/problems/majority-element-ii/#/description
#

def majority_element nums
  n = nums.length
  return nums if n < 2

  nums.sort!
  candidate1 = nums[0]
  count1 = 1
  candidate2 = nil
  count2 = 0
  i = 0

  while i < n
    if candidate1 != nums[i]
      candidate2 = nums[i]
      count = 1
      break
    end
    i += 1
  end

   puts "candidate1: #{candidate1}, candidate2: #{candidate2}"
  puts "count1: #{count1}, count2: #{count2}"

  i = 0
  while i < n
    if candidate1 != nums[i]
      count1 -= 1
    else
      count1 += 1
    end

    if candidate2 != nums[i]
      count2 -= 1
    else
      count2 += 1
    end

    candidate1 = nums[i] if count1 == 0
    candidate2 = nums[i] if count2 == 0

    i += 1
  end

  i = 0
  count1 = 0
  count2 = 0
  while i < n
    count1 += 1 if candidate1 == nums[i]
    count2 += 1 if candidate2 == nums[i]
    i += 1
  end

  puts "candidate1: #{candidate1}, candidate2: #{candidate2}"
  puts "count1: #{count1}, count2: #{count2}"
  result = []
  result << candidate1 if 3 * (count1.to_f / n) > 1
  result << candidate2 if 3 * (count2.to_f / n) > 1

  return result.uniq
end

arr = [2, 2, 1, 1, 1, 2, 2]
result =  majority_element arr

result.each {|i| print "#{i}, " }
