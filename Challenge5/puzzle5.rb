class Puzzle5
  attr_accessor :combustible_units, :input_string

  def initialize
    @combustible_units = generate_combustible_units
    @combustible_pairs = generate_combustible_pairs
    @input_string = set_input_string
  end

  def generate_combustible_units
    caps = ('A'..'Z').to_a
    lowers = ('a'..'z').to_a
    combustible_units = []
    caps.each_with_index do |letter, idx|
      first_string = letter + lowers[idx]
      second_string = lowers[idx] + letter
      combustible_units << first_string
      combustible_units << second_string
    end
    combustible_units
  end

  def generate_combustible_pairs
    i = 0
    pairs = []
    while i < combustible_units.length
      pairs << [combustible_units[i], combustible_units[i + 1]]
      i += 2
    end
    pairs
  end

  def set_input_string
    input_string = ''
    File.open('./input.txt').each do |line|
      input_string += line.chomp
    end
    input_string
  end

  def explode
    remaining = remove_combustibles(input_string, combustible_units)
    remaining.length
  end

  def find_smallest_combustible
    length_arr = []
    @combustible_pairs.each do |pair|
      new_string = input_string.gsub(pair[0][1], '')
      new_string.gsub!(pair[0][0], '')
      length_arr << remove_combustibles(new_string, combustible_units - pair).length
    end
    p length_arr
    p length_arr.min
  end

  def no_combustibles?(input)
    combustible_units.none? { |unit| input.include?(unit) }
  end

  def remove_combustibles(input_str, units)
    input = input_str
    idx = {}
    units.map { |unit| idx[unit] = 0 }

    until idx.empty?
      units.each do |combustible|
        index = input.index(combustible)
        if index.nil?
          idx.delete(combustible)
        else
          idx[combustible] = index
        end
      end

      index_to_remove = idx.min_by { |_k, v| v }
      if index_to_remove
        index_to_remove = index_to_remove[1]
        input = input[0...index_to_remove] + input[index_to_remove + 2..-1]
      end
    end
    input
  end
end

puz = Puzzle5.new
puz.find_smallest_combustible
