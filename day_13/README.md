# Day 13: Point of Incidence

In typical fashion, part 1 was not terribly difficult. For horizontal reflections I just looked for 2 adjacent rows that were the same and then compared rows above and below until either I reached an edge or found a difference. Vertical reflections were similarly tackled. In both instances, I abandoned the comparison as soon as I found a difference. I should have known....

Part 2 was _very_ intimidating at first. The thought of changing every item in a pattern, one at a time, and running a full check was annoying, and certainly not the intended solution.

I pondered a bit about some possible ways to tackle this. I knew I was going to have to look for pairs of rows/columns with just a single difference between them. I was overly-complicating things, though, by being stuck on the idea that I would have to know _which_ item needed to be changed, so that I could flip it and calculate the reflection.

Off to reddit for some inspiration. There I found an incredible suggestion. For part 1, we're looking for a reflection that has ZERO differences in it. But for part 2, we're looking for a reflection that has exactly ONE difference in it. We don't actually have to know which item needs to be flipped.

So, rather than looking for differences and stopping the moment I see one, I modified my code to count differences between each pair of rows/columns and keep a running tally. As soon as the diff tally exceeds my target (0 for part 1, 1 for part 2), I know the current adjacent pair has no reflection. If the reflection reaches the edge, I do one final check to make sure that the target number of diffs has been reached. This prevents an already existing clean reflection that has no diffs (found in part 1) from being used in part 2.
