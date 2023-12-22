
This was just about one of the dumbest puzzles ever. Rather than being a programming puzzle, it was a mathematical/input analysis puzzle. Dumb.

Part 1 was certainly easy enough - just a basic BFS for 64 steps, count the number of tiles hit on that step.

Part 2 was just stupid. Stupid, stupid, stupid. I thought about this for a while, then just jumped to reddit for the answer. And reddit was essentially just "fuck, this was stupid."

I tried following along with [this](https://github.com/villuna/aoc23/wiki/A-Geometric-solution-to-advent-of-code-2023,-day-21), but my resulting value is too low. Obviously there's an error in my code, but I'm just not interested in figuring it out.

Tired of wasting another single moment on this, I went to [this](https://github.com/mgtezak/Advent_of_Code/blob/master/2023/Day_21.py) to get some code to produce an answer. Thankfully it worked.

Got the star, and I fully admit to requiring outright theft to do it. Never let us speak of this again.

*UPDATE*: Ok, I couldn't let it sit. I went back through the math and realized my error. I need to add up all of the full even blocks, all of the full odd blocks, all of the even corners, and then substract all of the odd corners. My original math was: (all even blocks) + (all odd blocks) - (all odd corners + all even corners). ie, I was substracting all of the corners, instead of just the odd ones. Fixed it up and now my code gets the correct answer.
