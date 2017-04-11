#!/usr/bin/env ruby
#
s = gets.strip

array = s.split("")
count = 1

array.each {|i| count += 1 if ('A'..'Z').include? i }

puts count
