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
    i = 0;
    while i < 1100
      j = 0
      new_arr = []
      while j < 1100
        new_arr << 0
        j += 1
      end
      arr << new_arr
      i += 1
    end
    arr
  end

  def input_claims
    File.open('./input.txt').each do |line|
      parse_line(line)
    end
    find_all_overlap
    puts overlap
  end

  def parse_line(line)
    claim = Claim.new(line.split(' '))
    claims << claim
    add_to_fabric(claim)
  end

  def add_to_fabric(claim)
    (1..claim.width).each do |i|
      (1..claim.height).each do |j|
        row = claim.vertical_offset + j
        column = claim.horizontal_offset + i
        fabric[row][column] = fabric[row][column].zero? ? 1 : 2
      end
    end
  end

  def find_all_overlap
    fabric.each do |row|
      @overlap += row.select { |inch| inch == 2 }.count
    end
  end
end

puz = Puzzle3.new
puz.input_claims
