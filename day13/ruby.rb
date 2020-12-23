require 'prime'
require 'benchmark'

source = 'input.txt'

timestamp, schedule_line = File.readlines(source)
schedule = schedule_line.split(',').map { |x| x.chomp }

busses = schedule
  .each_with_index
  .map { |bus, offset| bus != 'x' ? [bus.to_i, offset] : nil }
  .compact

current_timestamp = 0

# This is unnecessary because all bus numbers are prime numbers, but this more
# general solution works for busses that are non-prime!
#
# In this exercise, the result is an array with one element: the first bus number:
# [7]
timestamp_factors = busses.shift.first.prime_division.map(&:first)

Benchmark.bm do |x|
  x.report do
    busses.each do |(bus, offset)|
      # The first bus will depart at t + 0 every t = bus * i. In fact, for any bus it's
      # true that once you find a valid t where t + offset = bus * i, the next
      # valid t for that bus is (t + bus).
      #
      # So you can skip all the ts inbetween.
      skip_factor = timestamp_factors.inject(&:*)

      # Find the next t for which [t + offset = bus * i].
      loop do
        minutes_to_departure = bus - (current_timestamp % bus)
        break if minutes_to_departure == offset % bus
        current_timestamp += skip_factor
      end

      # If all busses are prime numbers (they are in AoC), the prime-only
      # solution would be:
      #
      # timestamp_factors.push(bus)
      #
      # The general purpose solution looks like this:
      timestamp_factors.push(*bus.prime_division.map(&:first))
      timestamp_factors.uniq!
    end
  end
end

puts current_timestamp