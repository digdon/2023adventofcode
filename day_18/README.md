# Day 18: Lavaduct Lagoon

Part 1 was pretty easy. I worked out the trench points (corners) and then shifted them all so that the top-left would line up with 0,0 (instead of -5,6 or 15,3, for example). I drew out the trench into the grid, adding up the perimeter as I went along. From there, I did a simple flood fill, counting all of the cells that got filled. The result is the flood count plus the perimeter.

Of course, that wasn't going to work for part 2. Went to Reddit for inspiration and essentially [Pick's Theorem](https://en.wikipedia.org/wiki/Pick's_theorem) and the [shoelace formula](https://en.wikipedia.org/wiki/Shoelace_formula) were the flavours of the day. I understood the gist of the shoelace formula, but I kept ending up with a negative number. I was doing (x<sub>2</sub>\*y<sub>1</sub>)-(y<sub>2</sub>\*x<sub>1</sub>) instead of the more correct (x<sub>1</sub>\*y<sub>2</sub>)-(y<sub>1</sub>\*x<sub>2</sub>). This can be fixed by taking the absolute value.

I was then taking half of the area, per the formula, then trying to apply Pick's theorem, but I was messing all of that up. Again, Reddit to the rescue. Of note are the following two comments:
* https://www.reddit.com/r/adventofcode/comments/18l0qtr/comment/kdveugr/
* https://www.reddit.com/r/adventofcode/comments/18l0qtr/comment/kdvrqv8/

These helped me work out the details to correctly add half of the perimeter.

_UPDATE: I've spent several days thinking about why the equation had to be (A+b)/2+1. I knew from the links above that I was solving for i+b, rather than just i, but I needed to understand the mechanics behind it._

Let's see if I can fully explain what's going on here for part 2...

Using the test data from part 1 (to keep things simple), we have a perimeter of 38 blocks. If we count up the number of empty blocks inside the trench, we have 24. If we add all of that up, we have 38 + 24 = 62, which is the answer we're looking for. So how do we get there?

Pick's theorem is A = i + b/2 - 1, where i is the number of interior points and b is the number of boundary points which, in this case, is the perimeter/length of the trench. We are able to calculate two values, based on the input: the perimeter (simply by adding up the counts as we 'draw' out the trench) and the area, which we can find by using the shoelace formula. We know the perimeter: 38. That will be _b_. Shoelace gives us the area (_A_) of the trench: 42. We need to solve for _i_, so let's rearrange the equation -> i = A - b/2 + 1. Substituting values, we have i = 42 - 38/2 + 1 -> 42 - 19 + 1 -> 24.

So now we have the number of blocks that make up the the trench and the number of blocks inside the trench: 38 and 24. Now we just add 'em up. Nice and simple.

Still, what's with the i+b and (A+b)/2+1 stuff I originally had? Let's break that down too.

We know the perimeter already and we're solving for the number of interior points: i = A - b/2 + 1. However, the answer we actually need is the sum of the interior and perimeter. Rearranging, we have i + b/2 = A + 1. We already know that the raw area value from shoelace needs to be halved -> i + b/2 = A/2 + 1. However, we need i + b, not b/2. That means we're a b/2 short. We fix that by adding b/2 to both sides: i + b/2 + b/2 = A/2 + b/2 + 1. Simplifying, we have i + b = A/2 + b/2 + 1 -> i + b = (A+b)/2 + 1.