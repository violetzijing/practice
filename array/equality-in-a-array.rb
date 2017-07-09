#!/usr/bin/env ruby
#
# https://www.hackerrank.com/challenges/equality-in-a-array
#

length = gets.chomp
array = gets.chomp.split(" ")

i = 0
hash = {}
while i < array.length
  index = array[i]
  if hash[index] == nil
    hash[index] = 1
  else
    hash[index] += 1
  end
  i += 1
end

biggest_count = 0
hash.each do |key, value|
  biggest_count = value if value > biggest_count
end

puts length - biggest_count

