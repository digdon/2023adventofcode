# Day 6: Wait For It

After a few tricky days, day 6 was a refreshing break.

Of course, the immediate solution would be to iterate through each button push scenario (push = 0, push = 1, push = 2), then loop through the remaining time to calculate the possible distance - basically O(n^2). I wasn't happy with that idea, so I studied the data for a minute and noticed that the length of the button push multiplied by the remaining time actually gives the distance. ie, with a race time of 7 ms, we get the following table:

| Hold time | Remaining time | Distance |
|---|---|---|
| 0 | 7 | 0 |
| 1 | 6 | 6 |
| 2 | 5 | 10 |
| 3 | 4 | 12 |
| 4 | 3 | 12 |
| 5 | 2 | 10 |
| 6 | 1 | 6 |
| 7 | 0 | 0 |

I tried this for the second race and came up with the correct values. So, that means we can calculate the number of wins with a single iteration of all of the time values - O(n).

That extra time studying the data paid off. For part 2, the race time increased 1 million fold (eg, from 60 to 60 million). O(n^2) at that level would be expensive, but with the solution I implemented in part 1, it was just a single iteration through the 60 million values.

*Update:*
While looking at the table above, I realized that the second half of the table mirrors the first half. That means I really only need to work through the first half of the numbers and then double to get the number of wins.

But wait! Everything is a loss until the first winning value, at which point everything after that is a winner (considering just the first half). So, instead of counting winners, I just need to start at 0 and count losers until I hit the first winner. Double that, substract from the total number of scenarios, and that's the answer.
