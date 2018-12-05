require 'pry'
class Guard
  attr_accessor :id, :shifts, :description, :sleep_times

  def initialize(description)
    @description = description
    @sleep_times = {}
    parse_guard_description
  end

  def parse_guard_description
    @id = parse_id
    parse_date(description)
  end

  def parse_id
    # nests the result in [[RESULT]] so need to flatten and then take result
    @id = description.scan(/[#](\d*)/).flatten.first.to_i
  end

  def parse_date(description_string)
    date_description = description_string[1..10]
    time_description = description_string.scan(/\d{2}:\d{2}/).first
    generate_sleep_times_array(date_description)
    if description_string.include?('wakes up')
      sleep_char = 'w'
    elsif description_string.include?('falls asleep')
      sleep_char = 'a'
    elsif description_string.include?('#')
      time_description[3..4] = '00' if time_description[0..1] != '00'
      sleep_char = 's'
    end
    sleep_times[date_description][time_description[3..4].to_i] = sleep_char
  end

  def total_sleep
    @total_sleep ||= sleep_times.values.flatten.reduce(:+)
  end

  def determine_sleep
    sleep_times.each do |_k ,v|
      if v.include?('a') && v.include?('w') && (v.index('w').to_i < v.index('a').to_i)
        (0..v.index('w')).each { |i| v[i] = 1 }
      end
      asleep = false
      v.each_with_index do |num, idx|
        asleep = true if num == 'a'
        asleep = false if num == 'w'
        v[idx] = 1 if asleep && num != 1
      end
      v[v.index('w')] = 0 while v.index('w')
      v[v.index('s')] = 0 while v.index('s')
    end
  end

  def generate_sleep_times_array(date_description)
    sleep_times[date_description] ||= Array.new(60, 0)
  end

  def sleep_minutes_with_frequency
    occurence = Hash.new
    (0...60).each { |i| occurence[i] = 0 }
    sleep_times.values.each do |sleeps|
      sleeps.each_with_index do |time, idx|
        next unless time == 1

        occurence[idx] += sleeps[idx]
      end
    end
    occurence
  end

  def sleepiest_minute
    sleep_minutes_with_frequency.max_by{ |k, v| v }[0]
  end
end
