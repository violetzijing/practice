#!/usr/bin/env ruby
#

def generalizedGCD(num, arr)
    # WRITE YOUR CODE HERE
  result = arr[0]
  while i = 1; i < n; i++ do
    result = gcd(arr[i], result)
  end

  return result
end

def gcd(a, b)
    while a != b do
        if a > b
            a = a - b
        else
            b = b - a
        end
    end
    return a
end

puts gcd(60, 96)

