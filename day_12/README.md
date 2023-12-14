# Day 12: Hot Springs
I knew the brutal days were coming and this did not disappoint...

Part 1 wasn't terribly difficult and I tried a variety of techniques that all worked:
1. Using the group counts to build a regular expression, then building the spring values by replacing ? characters with . and #
   * I used some nifty bit twiddling to work out the ? assignments
1. That seemed slow, so then I just walked through the spring list, keeping track of everything as I went along (this will become important later). Every step involved a recursive call, with ? being handled in a special way
1. That also seemed inefficient. I figured that some sort of binary tree mechanism might be good so that I could kill off entire branches as soon as I noticed they would fail.

All of these worked for part 1, with solutions 2 and 3 being particularly effective. However, I was having difficulty keeping track of things, so I manipulated the spring data as I went along, and that seemed to make things easier. That was going to become a problem...

Part 2. Good times. I tried all kinds of variations to try to make this work, but I kept coming up empty-handed. Found some discussion on Reddit that talked about memoization. No idea what that's all about. Brought it up at work and got some clarity about how the concept works. Essentially, when you've calculated something, store the parameters that got you to that point and the result in a cache. That way, if you come across that scenario again, you can just look up the previously-calculated result.

I understood the concept, but generating the key (the scenario parameters) didn't make any sense - I didn't know really what to use. Thinking that my manipulated spring data was going to be a problem, I started from scratch and worked out a solution that just follows along the items, keeping track of everything as I go. I also went back to making a recursive call for every spring.

With the rewrite, it was obvious that what I needed for the key was the current spring position, the number of # in a row, and the current group. In hindsight, I was essentially doing this in most of my previous implementations, so I already had everything I needed....

Anyway, new, fresh, performant code, and I learned a new technique. I'll call all of that a win.
