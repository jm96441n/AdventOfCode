require_relative './guard.rb'

class Puzzle4
  attr_accessor :guards

  def initialize
    @guards = []
  end

  def find_sleepiest_guard
    parse_file
    @guards.map(&:determine_sleep)
    @sleepiest_guard = guards.max_by(&:total_sleep)
    puts @sleepiest_guard.id * @sleepiest_guard.sleepiest_minute
    sleepiest_minutes = {}
    guards.map { |g| sleepiest_minutes[g.id] = g.sleep_minutes_with_frequency.max_by{ |k, v| v }[1] }
    most_sleeping_guard_with_freq = sleepiest_minutes.max_by { |k, v| v }
    most_sleeping_guard = most_sleeping_guard_with_freq[0]
    frequency = most_sleeping_guard_with_freq[1]
    sleepiest_minutes = guards.find { |g| g.id == most_sleeping_guard }.sleep_minutes_with_frequency
    minute_mode = sleepiest_minutes.values.index(frequency)
    puts minute_mode * most_sleeping_guard
  end

  def parse_file
    lines =[]
    File.open('input.txt').each do |line|
      lines << line
    end
    lines.sort.each do |line|
      if line.include?('#')
        add_guard(line)
      else
        update_guard(line)
      end
    end
  end

  def find_guard_by_id(line)
    id = line.scan(/[#](\d*)/).flatten.first.to_i
    guards.find { |g| g.id == id }
  end

  def update_guard(line)
    @current_guard.description = line
    @current_guard.parse_date(line)
  end

  def add_guard(line)
    guard = find_guard_by_id(line)
    if guard
      guard.description = line
      guard.parse_date(line)
    else
      guard = Guard.new(line)
      guards << guard
    end
    @current_guard = guard
  end
end

puz = Puzzle4.new
puz.find_sleepiest_guard
