#!/usr/bin/env ruby
# https://www.hackerrank.com/challenges/ruby-case/problem
#

class Hacker
end
class Submission
end
class TestCase
end
class Contest
end
def identify_class(obj)
    # write your case control structure here
    case obj.name
    when "Hacker"
        puts "It's a Hacker!"
    when "Submission"
        puts "It's a Submission!"
    when "TestCase"
        puts "It's a TestCase!"
    when "Contest"
        puts "It's a Contest!"
    else
        puts "It's an unknown model"
    end
end

puts identify_class Hacker
puts identify_class Submission
puts identify_class TestCase
puts identify_class Contest
