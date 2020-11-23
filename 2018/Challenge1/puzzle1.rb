class Puzzle1
  attr_accessor :current_value, :frequency

  def initialize
    @frequency = {}
    @first_number_shown_twice = nil
    @current_value = 0
  end

  def find_frequency
    file_parse('1_input.txt') while @first_number_shown_twice.nil?
    puts "First frequency to be shown twice: #{@first_number_shown_twice}"
    puts "Current frequency #{@current_value}"
  end

  def file_parse(file_name)
    File.open("./#{file_name}").each do |line|
      change = parsed_number(line)
      if line_is_adding?(line)
        @current_value += change
      else
        @current_value -= change
      end
      assign_frequency
    end
  end

  def line_is_adding?(line)
    line[0] == '+'
  end

  def parsed_number(line)
    line[1..-1].to_i
  end

  def assign_frequency
    return unless @first_number_shown_twice.nil?

    if @frequency[@current_value]
      @frequency[@current_value] += 1
      @first_number_shown_twice = @current_value
    else
      @frequency[@current_value] = 1
    end
  end
end

puzzle = Puzzle1.new
puzzle.find_frequency
