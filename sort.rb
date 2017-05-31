#!/usr/bin/env ruby
#
#
require 'pry'
def print_array array
  array.each {|i| print "#{i}, " }
end

def bubble_sort array
  i = 0
  j = 0
  while i < array.length
    while j < array.length - 1
      array[j], array[j + 1] = array[j + 1], array[j] if array[j] > array[j + 1]
      j += 1
    end
    i += 1
  end
end

def insert_sort array
  length = array.length

  return if length == 1 or length == 0

  i = 1
  while i < length do
    tmp = array[i]
    j = i - 1
    while j >= 0 and tmp < array[j]
      array[j + 1] = array[j]
      j -= 1
    end
    array[j + 1] = tmp
    i += 1
  end
end

def merge left, right
  final = []

  until left.empty? or right.empty?
    final << (left.first < right.first ? left.shift : right.shift)
  end
 
  return final + left + right
end

def merge_sort array
  return array if array.length < 2

  pivot = array.length / 2
 
  array = merge(merge_sort(array[0...pivot]), merge_sort(array[pivot..-1]))

  return array
end

def quick_sort array
  return array if array.length < 2
  left, right = array[1..-1].partition {|y| y <= array.first }
  quick_sort(left) + array.first + quick_sort(right)
end

array = [1, 3, 2, 4, 7, 5]
#bubble_sort array
#insert_sort array
#puts merge_sort array
quick_sort array
array.each {|i| print "#{i}, " }
