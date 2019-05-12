#!/usr/bin/env ruby
#
def optimalUtilization(deviceCapacity, foregroundAppList, backgroundAppList)
	# WRITE YOUR CODE HERE
	results = []
	foregroundAppList.sort {|a,b| a[1] <=> b[1]}
	backgroundAppList.sort {|a,b| a[1] <=> b[1]}

	i = foregroundAppList.length - 1
	j = 0
	while foregroundAppList[i][1] >= deviceCapacity do
	    i -= 1
	end
    
	while i >= 0 && j < backgroundAppList.length do
    puts foregroundAppList[i][1]
    puts backgroundAppList[j][1]
	    if foregroundAppList[i][1] + backgroundAppList[j][1] > deviceCapacity
  	    i -= 1
  	    j += 1
	        next
	    end
	    results << [i, j]
	    i -= 1
	    j += 1
	end
	
	return results
end

fore = [[1, 8], [2, 7], [3, 14]]
back = [[1, 5], [2, 10], [3, 14]]
puts optimalUtilization(20, fore, back)
