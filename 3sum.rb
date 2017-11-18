#!/usr/bin/env ruby
#
# https://leetcode.com/problems/3sum/description/
#

def three_sum nums
  return [[0, 0, 0]] if nums.uniq == 0 and nums.length > 2
  return [] if nums.length < 3
  nums.sort!
  result = []

  m = 0
  while m < nums.length - 2
    i = m + 1
    j = nums.length - 1
 
    a = nums[m]
    while i < j
      b = nums[i]
      c = nums[j]

      if (a + b + c) == 0
        result << [a, b, c]
        j -= 1
        while j > i and nums[j] == nums[j + 1]
          j -= 1
        end
        i += 1
        while i < j and nums[i] == nums[i - 1]
          i += 1
        end
      elsif (a + b + c) > 0
        j -= 1
        while j > i and nums[j] == nums[j + 1]
          j -= 1
        end
      else
        i += 1
        while i < j and nums[i] == nums[i - 1]
          i += 1
        end
      end
    end
  
    m += 1
    while m < nums.length - 2 and nums[m] == nums[m - 1]
      m += 1
    end
  end

  return result
end

a = three_sum [-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]
a.each do |i|
  i.each {|n| print "#{n}, "}
  puts
end
