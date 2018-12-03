class Puzzle2
  attr_accessor :double_count, :triple_count

  def initialize
    @double_count = 0
    @triple_count = 0
    @checksum = 0
    @lines = []
  end

  def process_ids
    File.open('./input.txt').each do |line|
      find_repeats(line)
      build_lines_array(line)
    end
    correct_id
  end

  def correct_id
    correct_id = []
    @lines.each do |line|
      break unless correct_id.empty?

      (@lines - line).each do |oline|
        if line.zip(oline).count { |a, b| a != b } == 1
          differing_letter = line - oline
          correct_id = line - differing_letter
        end
      end
    end
    correct_id.join
  end

  def checksum
    @checksum = @double_count * @triple_count
  end

  private

  def build_lines_array(line)
    @lines << line.chomp.split('')
  end

  def find_repeats(id_string)
    letters = id_string.chomp.split('')
    uniq_letters = letters.uniq
    return if no_duplicates?(letters, uniq_letters)

    letters_hash = count_repeats(letters)

    values = letters_hash.values
    @double_count += 1 if values.include?(2)
    @triple_count += 1 if values.include?(3)
  end

  def count_repeats(letters)
    letters_hash = {}
    letters.each do |letter|
      if letters_hash[letter]
        letters_hash[letter] += 1
      else
        letters_hash[letter] = 1
      end
    end
    letters_hash
  end

  def no_duplicates?(letters, uniq_letters)
    letters == uniq_letters
  end
end

puz = Puzzle2.new
checksum = puz.checksum
p checksum
p puz.process_ids
