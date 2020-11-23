class Claim
  attr_accessor :height, :width, :horizontal_offset, :vertical_offset, :claim_number

  def initialize(chars)
    @height = parse_height(chars)
    @width = parse_width(chars)
    @horizontal_offset = parse_horizontal_offset(chars)
    @vertical_offset = parse_vertical_offset(chars)
    @claim_number = parse_claim(chars)
  end

  def parse_claim(line)
    line[0]
  end

  def parse_width(line)
    line[3].split('x')[0].to_i
  end

  def parse_height(line)
    line[3].split('x')[1].to_i
  end

  def parse_horizontal_offset(line)
    line[2].split(',')[0].to_i
  end

  def parse_vertical_offset(line)
    line[2].split(',')[1][0..-1].to_i
  end
end
