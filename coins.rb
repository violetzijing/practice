#!/usr/bin/env ruby
#

require 'pry'

def count_coins sum
  coins = []
  d = []
  d[0] = 0
  given_coins = [1, 3, 5]
  max = 2**62 - 1

  i = 1
  while i <= sum
    tmp = []
    tmp_coins = []
    j = 0
    while j < given_coins.length
      if i < given_coins[j]
        tmp[j] = max
      else
        tmp[j] = d[i - given_coins[j]] + 1
        tmp_coins[j] = given_coins[j]
      end
      j += 1
    end
    d[i] = tmp.min
    coins[i] = tmp_coins[tmp.index(d[i])]
    i += 1
  end

  i = sum
  choices = []
  while i > 0
    choices << coins[i]
    i = i - coins[i]
  end
  
  return choices
end

puts(count_coins(11))
