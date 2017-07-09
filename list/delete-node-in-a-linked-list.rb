#!/usr/bin/env ruby
#
# https://leetcode.com/problems/delete-node-in-a-linked-list/#/description
#

require_relative 'list-helper'

def delete_node node
  node.val = node.next.val
  node.next = node.next.next
  nil
end
