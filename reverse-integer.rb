#!/usr/bin/env ruby
# https://leetcode.com/submissions/detail/99878686/

def reverse x
  result = 0
  int_max1 = 2147483647
  int_max2 = 2147483648

  if x > 0
    loop do
      return 0 if result > int_max1 / 10
      result = result * 10 + x % 10
      x = x / 10
      break if x == 0
    end
  else
    x = 0 - x
    loop do
      return 0 if result > int_max2 / 10
      result = result * 10 + x % 10
      x = x / 10
      break if x == 0
    end

    result = 0 - result
  end

  return result
end



puts reverse -2147483412
