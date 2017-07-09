#!/usr/bin/env ruby
#
#

require 'pry'

def divide_stick total
  price = [0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30]
  d = []
  d[0] = 0

  # i is length
  # d[i] is total price

  i = 1
  while i <= total
    j = 1
    tmp = []
    tmp[0] = -1

    while j < price.length
      if i < j
        tmp[j] = -1
      else
        tmp[j] = d[i - j] + price[j]
      end
      j += 1
    end

    d[i] = tmp.max
    i += 1
  end
  
  return d[total]
end

puts divide_stick(4)
