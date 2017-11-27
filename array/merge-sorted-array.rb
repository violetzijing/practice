#!/usr/bin/env ruby
#
# https://leetcode.com/problems/merge-sorted-array/#/description
#

def merge nums1, m, nums2, n
  i = m - 1
  while i >= 0
    nums1[i + n] = nums1[i]
    i -= 1
  end

  i = n
  j = 0
  k = 0
  while i < n + m and j < n
    if nums1[i] > nums2[j]
      nums1[k] = nums2[j]
      j += 1
    else
      nums1[k] = nums1[i]
      i += 1
    end
    k += 1
  end

  if j < n
    while j < n
      nums1[j + m] = nums2[j]
      j += 1
    end
  end

  return nums1
end

nums1 = [1, 2, 3]
nums2 = [4]
result = merge nums1, nums1.length, nums2, nums2.length
result.each {|i| print "#{i}, " }
