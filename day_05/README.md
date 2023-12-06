# Day 5: If You Give A Seed A Fertilizer
Another one of those classic "part 1 is easy, part 2 looks like it's an extension of part 1, but it's not" days.

Part 1 was pretty simple - build some maps, then use the seed to look up a number in the soil map, use THAT value to look up a number in the fertilizer map, and so on. Lowest number wins.

Part 2 is essentially the same, except instead of a small set of seeds (eg,: 20), it's lists of huge ranges. I wrote up an initial solution to use the part 1 code to just brute force the solution. I was lucky in that my particular puzzle input took only 2 minutes to process. Good enough....

Except that it wasn't. I had some extra time, so I figured I'd see if I could do better.

Rather than trying to process each seed in those huge ranges, one by one, I knew I was going to have to process the seeds as ranges, similar to how I handled the huge ranges for part 1 for the simple lookups. Took some time to work out how to do it, but I eventually figured out how to calculate intersections of the seed range and conversion range (eg, seed to soil). When I find an intersection, I translate the intersection from the original values to the new converted values. For anything that falls outside of the intersection, I create new ranges and add then to the list of ranges to be processed. If a range has no intersection in a particular conversion map, we leave them as is.

The end result? The code (part 1 and part 2) runs in a total of just over 100ms. *That's* good enough...
