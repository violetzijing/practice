#!/usr/bin/env ruby
#
require 'pry'

class Node
  attr_accessor :left, :right, :key, :parent

  def initialize key, left = nil, right = nil, parent = nil
    @key = key
    @left  = left
    @right = right
    @parent = parent
  end
end

class Tree
  attr_accessor :root

  def initialize root = nil
    @root = root
    @level = {}
  end

  def invert_tree
    invert_tree_node root
  end

  def invert_tree_node node
    return unless node

    node.left, node.right = node.right, node.left
    invert_tree_node node.left
    invert_tree_node node.right
  end

  def insert key
    if @root == nil
      @root = Node.new key
      return
    end

    node = @root
    parent = nil

    while node != nil do
      parent = node
      if key < node.key
        # left
        node = node.left
      elsif key > node.key
        # right
        node = node.right
      else
        return
      end
    end

    if key > parent.key
      parent.right = Node.new key
      parent.right.parent = parent
    else
      parent.left = Node.new key
      parent.left.parent = parent
    end
  end

  def get_node key
    node = @root
    parent = nil

    while node != nil do
      parent = node
     
      if key > node.key
        # right
        node = node.right
      elsif key < node.key
        # left
        node = node.left
      else
        return node
      end
    end
  end

  def inorder_traverse_driver
    self.inorder_traverse (@root)
  end

  def inorder_traverse node
    inorder_traverse node.left if node.left != nil
    puts node.key
    inorder_traverse node.right if node.right != nil
  end

  def pre_order_traverse_driver
    self.pre_order_traverse @root
    puts
  end

  def pre_order_traverse node
    if node.left == nil and node.right == nil
      flag = true
    else
      flag = false
    end

    print node.key
    print "(" unless flag
    pre_order_traverse node.left if node.left != nil
    print ", " unless flag
    pre_order_traverse node.right if node.right != nil
    print ")" unless flag
  end

  def left_rotation node
    return if node.right == nil
  
    pivot = node
    new_top = pivot.right
    if @root == pivot
      @root = new_top
    else
      top_parent = pivot.parent
  
      if top_parent.left == pivot
        top_parent.left = new_top
      elsif top_parent.right == pivot
        top_parent.right = new_top
      end
    end
  
    pivot.right = new_top.left
    new_top.left = pivot
    pivot_parent = pivot.parent
    pivot.parent = new_top
    new_top.parent = pivot_parent
    pivot.right.parent = pivot if pivot.right != nil
  end
  
  def right_rotation pivot
    return if pivot.left == nil

    new_top = pivot.left

    if pivot == @root
      @root = new_top
    else
      top_parent = pivot.parent

      if top_parent.left == pivot
        top_parent.left = new_top
      elsif top_parent.right == pivot
        top_parent.right = new_top
      end
    end

    pivot.left = new_top.right
    new_top.right = pivot
    pivot_parent = pivot.parent
    pivot.parent = new_top
    new_top.parent = pivot_parent
    pivot.left.parent = pivot if pivot.left != nil
  end

  def delete_node key
    node = get_node key
    child_count = 0
    child_count += 1 if node.left != nil
    child_count += 1 if node.right != nil

    if child_count == 0
      if node.parent != nil
        node.parent.left = nil if node.parent.left == node
        node.parent.right = nil if node.parent.right == node
      else
        @root = nil
      end
    elsif child_count == 1
      if node.parent != nil
        if node.parent.left == node
          node.parent.left = node.left if node.left != nil
          node.parent.left = node.right if node.right != nil
        else
          node.parent.right = node.left if node.left != nil
          node.parent.right = node.right if node.right != nil
        end
      else
        @root = node.left if node.left != nil
        @root = node.right if node.right != nil
      end
    elsif child_count == 2
      successor = self.successor node
      node.key = successor.key
      successor.parent.left = nil if successor.parent.left == successor
      successor.parent.right = nil if successor.parent.right == successor
      successor.parent = nil
    end
  end

  def delete_node2 key
    z = get_node key

    if z.left == nil or z.right == nil
      y = z
    else
      y = self.successor z
    end

    if y.left != nil
      x = y.left
    else
      x = y.right
    end

    x.parent = y.parent if x != nil

    if y.parent == nil
      @root = x
    else
      if y == y.parent.left
        y.parent.left = x
      else
        y.parent.right = x
      end
    end

    if y != z
      z.key = y.key
    end
  end

  def successor node
    return min_value node.right if node.right != nil

    parent = node.parent
    while parent != nil do
      break if node != parent.right

      node = parent
      parent = parent.parent
    end

    return parent
  end

  def min_value node
    current = node
    
    while current.left != nil do
      break if current.left == nil
      current = current.left
    end

    return current
  end
end


tree = Tree.new
items = [20, 15, 10, 25, 18, 12, 30]
items.each {|i| tree.insert i }
#tree.left_rotation node
tree.pre_order_traverse_driver
tree.delete_node2 30
tree.pre_order_traverse_driver
