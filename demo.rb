#!/usr/bin/env ruby
#
require 'pry'

def solution s
  array = s.split("\n")
  
  total_time = {}
  total_price = {}


  array.each do |line|
    time = line.split(",")[0]
    num = line.split(",")[1]
    
    if total_time[num] == nil
      total_time[num] = time_to_second(time)
    else
      total_time[num] += time_to_second(time)
    end
    
    price = every_single_price time

    if total_price[num] == nil
      total_price[num] = price
    else
      total_price[num] += price
    end
  end

  max_total_time = total_time.values.max
  max_total_time_num = total_time.reject{|k,v|v != max_total_time }.keys

  tmp = []
  max_total_time_num.each do |num|
    tmp << total_price[num]
  end
  max_time_min_price_num = tmp.min

  return total_price.values.reduce(:+) - max_time_min_price_num
end

def every_single_price duration
  time = duration.split(":")

  price = 0
  if time[0].to_i == 0 and time[1].to_i < 5
    price = time_to_second(duration) * 3
  else
    total_time = time[0].to_i * 60 + time[1].to_i
    total_time += 1 if time[2] != "00"
    price = total_time * 150
  end
end

def time_to_second duration
  time = duration.split(":")

  total_second = time[0].to_i * 3600 + time[1].to_i * 60 + time[2].to_i
end


s = "00:01:07,400-234-090
     00:05:01,701-080-080
     00:05:00,400-234-090"

puts (solution s)
