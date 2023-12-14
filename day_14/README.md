# Day 14: Parabolic Reflector Dish

Another interesting one that involves looking for shortcuts.

Part 1 was very easy. I knew part 2 would involve tilting in all 4 directions, so I had things set up early to account for that.

Part 2, of course, was not so easy. Tilting the board 4 billion times just isn't feasible. It occurred to me that there would be a loop. I wasn't thinking clearly and thought that the loop would eventually repeat every 4 cycles. Of course, what I was really thinking was 4 tilts, meaning 1 cycle. Anyway, for whatever reason, I set a target of 1000 cycles and, oddly enough, it returned the correct answer. In disbelief, I went to reddit to try to figure out why using a value 1 millionth of the target worked. All of the solutions involved detecting a cycle, working out the length, then skipping ahead to get as close to the end as possible, and then working through the remaining cycles until the target count was reached.

I took some inspiration from a colleague for caching each cycle - sort all of the rocks into a predictable order, generate a string, and use that as the key. The strings are very long, so I took it a step further and generated a SHA256 hash from my Point string values.

It's very interesting that my puzzle data took 152 cycles to detect a loop of 14 cycles (138, 152, 166, etc).

Overall, a very interesting day and I learned some interesting things in the process.
