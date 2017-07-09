#!/usr/bin/env ruby
#
require 'colorize'
require 'pry'

class Node
  attr_accessor :left, :right, :key, :parent, :color, :value
  def initialize(key, left = nil, right = nil, parent = nil, color = "red", value = nil)
    @left = left
    @right = right
    @key = key
    @parent = parent
    @color = color
    @value = value
  end
end

class Tree
  attr_accessor :root
  def initialize root = nil
    @root = root
  end

  def insert key
    if @root == nil
      @root = Node.new key
      insert_fix @root
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

    added_node = Node.new key

    if key > parent.key
      parent.right = added_node
      parent.right.parent = parent
    else
      parent.left = added_node
      parent.left.parent = parent
    end

    insert_fix added_node
  end

  def insert_fix node
    if node == @root
      node.color = "black"
      return
    end

    return if node.parent.color == "black"

    node.color = "red"

    while node.parent != nil and node.parent.color == "red"
      if node.parent == node.parent.parent.left
        uncle = node.parent.parent.right
        if uncle.color == "red" and node.parent.color == "red"
          # case 1
          node.parent.color = "black"
          uncle.color = "black"
          node.parent.parent.color = "red"
          node = node.parent.parent
        elsif node = node.parent.right
          # case 2, left rotate
          node = node.parent
          left_rotation node
          # case 3, right rotate
          node.parent.color = "black"
          node.parent.parent.color = "red"
          right_rotation node.parent.parent
        end
      else
        uncle = node.parent.parent.left
        if uncle.color == "red" and node.parent.color == "red"
          # case 1

          node.parent.color = "black"
          uncle.color = "black"
          node.parent.parent.color = "red"
          node = node.parent.parent
        elsif node = node.parent.left
          # case 3, left rotate
          node = node.parent
          right_rotation node
          # case 2, right rotate
          node.parent.color = "black"
          node.parent.parent.color = "red"
          left_rotation node.parent.parent
        end
      end
    end

    @root.color = "black"
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

   def inorder_traverse_driver
    self.inorder_traverse (@root)
  end

  def inorder_traverse node
    inorder_traverse node.left if node.left != nil
    puts node.key.to_s.red if node.color == "red"
    puts node.key.to_s.blue if node.color == "black"
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

    print node.key.to_s.red if node.color == "red"
    print node.key.to_s.blue if node.color == "black"
    print "(" unless flag
    pre_order_traverse node.left if node.left != nil
    print ", " unless flag
    pre_order_traverse node.right if node.right != nil
    print ")" unless flag
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

  def delete_node key
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

    binding.pry
    x.parent = y.parent

    if y.parent == nil
      @root = x
    else
      if y == y.parent.left
        y.parent.left = x
      else
        y.parent.right = x
      end
    end

    z.key = y.key if y != z

    delete_node_fixup x if y.color == "black"
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

    z.key = y.key if y != z
   
    binding.pry
    delete_node_fixup x if y.color == "black"
  end

  def delete_node_fixup x
    binding.pry
    while x != @root and x.color == "black"
      if x == x.parent.left
        w = x.parent.right
        if w.color == "red"
          # case 1
          x.parent.color = "red"
          self.left_rotation x.parent
          w = x.parent.right
        end

        if w.left.color == "black" and w.right.color == "black"
          # case 2
          w.color = "red"
          x = x.parent
        else
          # case 3
          if w.right.color == "black"
            w.left.color = "black"
            w.color = "red"
            self.right_rotation w
            w = x.parent.right
          end
          # case 4
          w.color = x.parent.color
          x.parent.color = "black"
          w.right.color = "black"
          self.left_rotation x.parent
          x = @root
        end
      else  
        w = x.parent.left
        if w.color  == "red"
          # case 1
          x.parent.color = "red"
          self.right_rotation x.parent
          w = x.parent.left
        end
        
        if w.right.color == "black" and w.left.color == "black"
          # case 2
          w.color = "red"
          x = x.parent
        else
          # case 3
          if w.left.color == "black"
            w.right.color = "black"
            w.color = "red"
            self.left_rotation w
            w = x.parent.left
          end
          # case 4
          w.color = x.parent.color
          x.parent.color = "black"
          w.right.color = "black"
          self.right_rotation x.parent
          x = @root
        end
      end
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

black = "black"
red = "red"
items = [11, 2, 14, 1, 7, 15, 5, 8, 4]
tree = Tree.new

items.each {|i| tree.insert i }

tree.pre_order_traverse_driver
tree.delete_node2 1
tree.pre_order_traverse_driver

=begin
node11 = Node.new(11)
node2 = Node.new(2)
node14 = Node.new(14)
node1 = Node.new(1)
node7 = Node.new 7
node15 = Node.new 15
node5 = Node.new 5
node8 = Node.new 8

tree.root = node11

node11.color = black
node11.left = node2
node11.right = node14

node2.color = red
node2.parent = node11
node2.left = node1
node2.right = node7

node14.color = black
node14.parent = node11
node14.right = node15

node1.color = black
node1.parent = node2

node7.color = black
node7.parent = node2
node7.left = node5
node7.right = node8

node15.color = red
node15.parent = node14

node5.color = red
node5.parent = node7

node8.color = red
node8.parent = node7
=end

