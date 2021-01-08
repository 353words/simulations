### Introduction

> If you can write a for-loop, you can do statistics.
> - Jake Vanderplas

A lot of developers shy away from problems which involve statistics or probability. Which is shameful since in today's data-rich environment, you can gain a lot of insights from data.

In this blog post, I'll show you how to write a simulation tool which requires no knowledge in statistics or probability. Simulations are easy to write and can be a very effective tool in research. You only need some basic programming skills and a random number generator.

### Catan 

![](catan.png)


In the game of [Catan](https://en.wikipedia.org/wiki/Catan), you gain resources when a roll of two dice match a number of a tile on the board. At the beginning of the game, you place your settlements next to tiles with the idea of selecting titles that have a higher probability of being matched.

The first thing we need for our simulation is the ability to roll a die and get the result.

**Listing 1: Dice roll**  
```
09 // diceRoll simulate a dice roll.
10 func diceRoll() int {
11     // Intn(6) returns values in the range 0-5 (inclusive), we want 1-6.
12     return rand.Intn(6) + 1
13 }
```

Listing 1 shows how we will perform a single dice roll for our simulation. On line 12, we use the random `Intn` function from the `math` package and adjust the values returned to be 1-6 instead of 0-5.

Next, we need to simulate rolling two dice over and over to learn what sum’s of the two dice tend to be rolled the most.

**Listing 2: Dice roll simulation**  
```
15 // simulate executes “runs” two game cube rolls.
16 // It returns the percentage for each total of first and second roll.
17 func simulate(runs int) map[int]float64 {
18     counts := make(map[int]int)
19     for i := 0; i < runs; i++ {
20         val := diceRoll() + diceRoll()
21         counts[val]++
22     }
23 
24     // Convert from counts to fractions.
25     fracs := make(map[int]float64)
26     for val, count := range counts {
27         frac := float64(count) / float64(runs)
28         fracs[val] = frac
29     }
30 
31     return fracs
32 }
```

Listing 2 shows our `simulation` function. This function rolls two dice `runs` number of times and returns the percent of time any given sum was rolled over `n`. On line 18, we initialize a `counts` map that will hold the counts for the different sum’s rolled. On line 20, we roll the two dice and calculate the sum. On line 21, we update the count for that sum in the map. On lines 25-29, we convert the sums to fractions based on the total number of rolls and finally on line 31, we return the results.

Finally, we need a little application to run the simulation.

**Listing 3: Running the simulation**  
```
34 func main() {
35     fracs := simulate(1_000_000)
36     for i := 2; i <= 12; i++ {
37         fmt.Printf("%2d -> %.2f\n", i, fracs[i])
38     }
39 }
```

Listing 3 shows how we run the simulation. On line 35, we run the simulation and on lines 36-39, we print the results.

**Listing 4: Output**  
```
 2 -> 0.03
 3 -> 0.06
 4 -> 0.08
 5 -> 0.11
 6 -> 0.14
 7 -> 0.17
 8 -> 0.14
 9 -> 0.11
10 -> 0.08
11 -> 0.06
12 -> 0.03
```

Listing 4 shows the output of running the simulation one time. In this run, the sum of 7 is the winner with 6 and 8 in a tie for second place. Since Catan doesn’t use the number 7, this simulation is suggesting you should pick 6 or 8.

_Note: Running this simulation will return the same results every time. This is due to the fact that random number generators rely on an initial seed to generate pseudo-random numbers. If you want to see different results, the common practice is to seed the random number generator with the current time: rand.Seed(time.Now().Unix())_

### The Birthday Problem

The [birthday problem](https://en.wikipedia.org/wiki/Birthday_problem) asks what is the probability that in a group of people, at least two people will have the same birthday? Let's create a simulation for this problem that can answer this question for a group of 23 people. 

First, let’s write a function that randomly chooses a birthday for an individual who belongs to a given size group of people. If at any time two people end up with the same birthday, the function will immediately return true.

**Listing 5: Checking for same birthday in a random group**  
```
10 // simulateBirthdayMatches returns true if the same number is selected
11 // twice by the random number generator selecting a number between
12 // 0 and 365 for a specified group of people.
13 func simulateBirthdayMatches(numOfPeople int) bool {
14     const daysInYear = 365
15 
16     seen := make(map[int]bool)
17     for i := 0; i < numOfPeople; i++ {
18         day := rand.Intn(daysInYear)
19         if seen[day] {
20             return true
21         }
22         seen[day] = true
23     }
24 
25     return false
26 }
```

Listing 5 shows the `simulateBirthdayMatches` function. Based on a specified group of people, it will draw random birthdays. Birthdays are represented as days since the beginning of the year: we have 365 of these.

On line 16, we initialize a map for already seen birthdays. On line 18, we draw a random day and on line 19, we check if we've seen this day already. If that day was already seen, we return `true` on line 20. Otherwise we update the `seen` map on line 22. Finally, if there are no duplicated days, we return `false` on line 25.

Next, let’s write a function that can run the birthday match simulation over and over for a given size group of people.

**Listing 6: Birthday simulation**   
```
28 // simulateBirthdays returns the fraction of groups
29 // that have two people with the same birthday.
30 func simulateBirthdays(numOfPeople, runs int) float64 {
31     same := 0
32     for i := 0; i < runs; i++ {
33         if simulateBirthdayMatches(numOfPeople) {
34             same++
35         }
36     }
37 
38     return float64(same) / float64(runs)
39 }
```

Listing 6 shows the simulation code. We pass the number of people and the number of times we want to run the simulation. On line 31, we initialize the number of groups that had at least one duplicate birthday. On line 32, we check `runs` random groups and on line 33, we check if there was a duplicate birthday in the current group. If there is, we update the `same` counter on line 34. Finally on line 38, we return the fraction of groups that had the same birthday.

Finally, let’s write a small application to run the simulation.

**Listing 7: Running the simulation**  
```
41 func main() {
42     fmt.Println(simulateBirthdays(23, 1_000_00))
43 }
```

Listing 7 shows how we run the simulation. On line 41, we run the simulation and print the results.

**Listing 8: Output**  
```
0.50536
```

Listing 8 shows the output of our simulation. We see that the chance that two people in a group of 23 having the same birthday is about 50%, which matches the [expected result](https://en.wikipedia.org/wiki/Birthday_problem).

### Sick or Not?

The following question is taken from Nassim Taleb's [Fooled By Randomness](https://www.amazon.com/Fooled-Randomness-Hidden-Markets-Incerto/dp/0812975219) book which I highly recommend.

The book makes this statement which we will base this simulation on.

_The test of a disease presents a rate of 5% false positives. The disease strikes 1/1000 of the population. People are tested at random, regardless of whether they are suspected of having the disease. A patient’s test is positive. What is the probability of the patient being stricken with the disease?_

To be clear, a “false positive” is when a healthy person is diagnosed as sick.

**Listing 9: A random event**  
```
11 // oneChanceIn returns true if the number 1 is selected randomly
12 // from a group of n numbers.
13 func oneChanceIn(n int) bool {
14     return rand.Intn(n) == 1
15 }
```

Listing 9 shows the `oneChanceIn` function which returns `true` if the number 1 is selected randomly from a range of numbers. On line 13, we generate a random number and compare the number to 1. The probability of choosing the number 1 randomly (or any other number between 0 and n-1) is 1/n.


**Listing 10: The isSick**  
```
16 // isSick returns true if a randomly sampled person is sick.
17 func isSick() bool {
18     // The disease strikes 1/1000 of the population.
19     return oneChanceIn(1000)
20 }
``` 

Listing 10 shows the `isSick` function which returns `true` if a random sample from the population is sick or not. According to the problem statement, “The disease strikes 1/1000 of the population.”, so we use `oneChanceIn(1000)` on line 19.

**Listing 11: diagnosed**  
```

22 // diagnosed returns true if a person is sick or misdiagnosed as sick.
23 func diagnosed(sick bool) bool {
24     if sick {
25         return true // We're 100% correct in sick people.
26     }
27 
28     // The test of a disease presents a rate of 5% (1 in 20) false positives.
29     // (false positive = healthy diagnosed as sick)
30     return oneChanceIn(20)
31 }
```

Listing 11 shows the `diagnosed` function that simulates testing a person for the disease. On line 25, we return `true` if the person is sick since we assume the test is accurate 100% of the time for sick people. If the person is healthy, then on line 30 we have a 1 in 20 (5%) chance of incorrectly diagnosing the patient as sick (false positive).


**Listing 12: The simulation**  
```
33 // simulate run selects sampleSize random people and return the fraction of people
34 // actually sick from the total number of people diagnosed as sick.
35 func simulate(sampleSize int) float64 {
36     var numSick, numDiagnosed int
37 
38     for i := 0; i < sampleSize; i++ {
39         sick := isSick()
40         if sick {
41             numSick++
42         }
43 
44         if diagnosed(sick) {
45             numDiagnosed++
46         }
47     }
48 
49     return float64(numSick) / float64(numDiagnosed)
50 }
```

Listing 12 shows the simulation code. On line 36, we initialize the number of actual sick people (`numSick`) and the number of people diagnosed as sick (`numDiagnosed`). On line 38, we sample `sampleSize` people from the population and on line 39, we determine if the person is actually sick or not. On line 41, we increment the number of actual sick people if the person is actually sick and on line 45, we increment the number of people diagnosed as sick according to the output of `diagnosed`. Finally, on line 49 we return the ratio of people who are actually sick from the people who were diagnosed as sick.

**Listing 13: Running the simulation**  
```
52 func main() {
53     fmt.Println(simulate(1_000_000))
54 }
```

Listing 13 shows a run of the simulation. One line 53, we run the simulation for 1 million people and print the results.

**Listing 14: Simulation output**  
```
0.020525227299332736
```

Listing 14 shows the output of the simulation run. The chance that a person is sick given a positive test is about 2% which matches the [expected result](https://psychscenehub.com/psychinsights/well-understand-probabilities-medicine/).

### Conclusion

Simulations are a simple and powerful tool. You don't need to know advanced probability and statistics in order to solve data driven problems. All you need is `math/rand` and basic programming skills. After you get your results, you are encouraged to validate them. If you don't have the skills - ask around (I ask my friend & college [Shlomo Yona](https://www.mathematic.ai/) whenever I need some math guidance).

If you want to learn more, I recommend watching [Statistics for Hackers](https://www.youtube.com/watch?v=Iq9DzN6mvYA). The examples are in Python, but are very easy to follow. You should also read about the [Monte Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method) on Wikipedia and see the wide variety of applications it has.

You can find the code for these simulations [here](https://github.com/353words/simulations).
Catan photo by [Galen Crout](https://unsplash.com/@galen_crout)
