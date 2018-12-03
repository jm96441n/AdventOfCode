require_relative './claim.rb'

class Puzzle3
  attr_accessor :claims, :fabric, :overlap

  def initialize
    @fabric = generate_fabric
    @overlap = 0
    @claims = []
  end

  def generate_fabric
    arr = []
    (1..1100).each do |_|
      new_arr = []
      (1..1100).each { |_| new_arr << 0 }
      arr << new_arr
    end
    arr
  end

  def input_claims
    File.open('./input.txt').each do |line|
      parse_line(line)
    end
    find_all_overlap
    puts overlap
    claim = find_claim_with_no_overlaps
    puts claim.claim_number
  end

  def parse_line(line)
    claim = Claim.new(line.split(' '))
    claims << claim
    add_to_fabric(claim)
  end

  def add_to_fabric(claim)
    fabric_iterator(claim) do |i, j|
      row = claim.vertical_offset + j
      column = claim.horizontal_offset + i
      fabric[row][column] = fabric[row][column].zero? ? 1 : 2
    end
  end

  def find_all_overlap
    fabric.each do |row|
      @overlap += row.select { |inch| inch == 2 }.count
    end
  end

  def find_claim_with_no_overlaps
    current_claim = ''
    claims.each do |claim|
      current_claim = claim
      return current_claim if no_overlaps?(claim)
    end
  end

  def no_overlaps?(claim)
    fabric_iterator(claim) do |i, j|
      row = claim.vertical_offset + j
      column = claim.horizontal_offset + i
      return false if fabric[row][column] == 2
    end
    true
  end

  def fabric_iterator(claim)
    (1..claim.width).each do |i|
      (1..claim.height).each do |j|
        yield(i, j)
      end
    end
  end
end

puz = Puzzle3.new
puz.input_claims
