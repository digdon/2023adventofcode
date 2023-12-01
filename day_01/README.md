# Day 1: Trebuchet?!

First day of using Go. Not knowing the standard library makes things harder for me, but that's the point of this exercise.

Part 1 was pretty easy - a simple regex to pull out all of the digits, grab the first and last, build the number.

Part 2 was more complicated because of the number words (one, two, etc). I originally had a regex to solve this, and while it worked with the test/example data, it failed on the puzzle input on stuff like this: `three98oneightzn`. That oneight is a problem - is it one or is it eight? It turns out that, in this particular case, it's eight, so parsing this should produce `three 9 8 eight`, which will result in a value of 38 for the line. I took it a step further in allowing overlaps, resulting in `three 9 8 one eight`.

After coming up with the solution for part 2, I modified the part 1 code to use the same code. I've left the original solution in place, but commented out.
