#!/usr/bin/env ruby
def less(a, b) 
    i = 0
    while a[i] != " " do
        i += 1
    end
    j = 0
    while b[j] != " " do
        j += 1
    end
    i += 1
    j += 1
    while i < a.length && j < b.length do
        return -1 if a[i] == " " && b[j] != " "
        return 1 if a[i] != " " && b[j] == " "
        return -1 if a[i] < b[j] && a[i] != " " && b[j] != " "
        return 1 if a[i] > b[j]
        i += 1
        j += 1
    end
    if i == a.length && j != b.length
        return -1
    elsif i != a.length && j == b.length
        return 1
    end
    a_id = a.split(" ")[0]
    b_id = b.split(" ")[0]
    if a < b
        return -1
    else
        return 1
    end
end

puts less("125 r1", "10a echo show")
