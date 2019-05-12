#!/usr/bin/env ruby
#

def cellCompete(states, days)
	# Convert states to a number
	num = 0
	for i in 0..7
		num |= (states[i] << (7 - i))
	end
	while days > 0 do
		num ^= (num << 2)
		num >>= 1
		num &= 0xff
		days -= 1
	end
	res = []
	# Convert num to an array
	for i in 0..7 do
		# Get the last bit
		res.unshift(num & 1)
		num >>= 1
	end
	return res
end

states = [1, 1, 1, 0, 1, 1, 1, 1]
days = 2

puts(cellCompete(states, days))
