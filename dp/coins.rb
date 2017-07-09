#!/usr/bin/env ruby
#

def coin sum
  given_coins = [3, 5, 7]

  coin = []
  d = []
  d[0] = 0
  i = 1
  max = 2**62 - 1
  
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
    coin[i] = tmp_coins[tmp.index(d[i])]
    i += 1
  end

  return d[sum]
end

puts coin 11
