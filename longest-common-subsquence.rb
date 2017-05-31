#!/usr/bin/env ruby
#

require 'pry'

def print_b b
  for i in (0...b.length)
    for j in (0...b[i].length)
      print "#{b[j][i]} "
    end
    puts ""
  end
end

def lcs x, y
  return if x.length == 0 or y.length == 0

  # 1 presents up
  # 2 presents turning left
  # 3 presents ↖
  
  z = Array.new(x.length + 1) { Array.new(y.length + 1, 0) }
  b = Array.new(x.length + 1) { Array.new(y.length + 1, 0) }

  i = 1

  while i <= x.length
    j = 1
    while j <= y.length
      if x[i - 1] == y[j - 1]
        z[i][j] = z[i - 1][j - 1] + 1
        b[i][j] = "↖" 
      else
        if z[i - 1][j] > z[i][j - 1]
          z[i][j] = z[i - 1][j]
          b[i][j] = "↑"
        else
          z[i][j] = z[i][j - 1]
          b[i][j] = "←"
        end
      end

      j += 1
    end

    i += 1
  end
  
  print_b b

  return z[x.length][y.length]
end

def print_lcs b
  i = 0
  while i < b.length
    j = 0
    while j < b[i].length
      if b[i][j] == 1
        # 1 presents up
      elsif b[i][j] == 2
        # 2 presents left
      elsif b[i][j] == 3
        # 3 presents ↖
      end
    end
    i += 1
  end
end

x = "ABCBDAB"
y = "BDCABA"

puts lcs(x, y)


