# Day 18: Lavaduct Lagoon

Part 1 was pretty easy. I worked out the trench points (corners) and then shifted them all so that the top-left would line up with 0,0 (instead of -5,6 or 15,3, for example). I drew out the trench into the grid, adding up the perimeter as I went along. From there, I did a simple flood fill, counting all of the cells that got filled. The result is the flood count plus the perimeter.

Of course, that wasn't going to work for part 2. Went to Reddit for inspiration and essentially [Pick's Theorem](https://en.wikipedia.org/wiki/Pick's_theorem) and the [shoelace formula](https://en.wikipedia.org/wiki/Shoelace_formula) were the flavours of the day. I understood the gist of the shoelace formula, but I kept ending up with a negative number. I was doing (x<sub>2</sub>\*y<sub>1</sub>)-(y<sub>2</sub>\*x<sub>1</sub>) instead of the more correct (x<sub>1</sub>\*y<sub>2</sub>)-(y<sub>1</sub>\*x<sub>2</sub>). This can be fixed by taking the absolute value.

I was then taking half of the area, per the formula, then trying to apply Pick's theorem, but I was messing all of that up. Again, Reddit to the rescue. Of note are the following two comments:
* https://www.reddit.com/r/adventofcode/comments/18l0qtr/comment/kdveugr/
* https://www.reddit.com/r/adventofcode/comments/18l0qtr/comment/kdvrqv8/

These helped me work out the details to correctly add half of the perimeter.
