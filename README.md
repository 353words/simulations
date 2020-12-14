# Using Simulations

> If you can write a for-loop, you can do statistics
> - Jake Vanderplas

A lot of developers shy away from problems which involve statistics or probability. Which is shameful since in today's data-rich environment, you can gain a lot of insights from data.

In the blog post I'll show you a tool that requires no knowledge in statistics or probably - a simulation. Simulations are easy to write and can be a very effective tool in research. All you need are some programming skills and a random number generator.

## Catan 

![](catan.png)


In the game of [Catan](https://en.wikipedia.org/wiki/Catan), you gain resources if a roll of two dice matches a number of a tile. At the beginning of the game you place your settlements next to some tiles and would like to pick tiles that have higher probability of being matched.

**Listing 1: Dice roll**

```
9 // diceRoll simulate a dice roll
10 func diceRoll() int {
11     // Intn returns values in the range [0-6), we want [1-6]
12     return rand.Intn(6) + 1
13 }
```

Listing one shows one dice roll. On line 12 we adjust the values returned from `rand.Intn(6)` to be 1-6 instead of 0-5.

**Listing 2: Dice roll simulation**

```
15 // simulate run n simulation of two game cube rolls and returns the percentage for each total of first and second roll
16 func simulate(n int) map[int]float64 {
17     counts := make(map[int]int)
18     for i := 0; i < n; i++ {
19         val := diceRoll() + diceRoll()
20         counts[val]++
21     }
22 
23     // Convert counts to fractions
24     fracs := make(map[int]float64)
25     for val, count := range counts {
26         frac := float64(count) / float64(n)
27         fracs[val] = frac
28     }
29     return fracs
30 }
```

Listing 2 shows a simulation of `n` rolls of two dice. On line 17 we initialize a `counts` map that will count how many times each sum of two dice rolls we saw. On line 19 we simulate a roll of two dice and one line 20 we update the counts.
On lines 23-28 we convert the counts to fractions of the total amount of runs (`n`) and finally on line 29 we return the fractions.

**Listing 3: Running the simulation**

```
32 func main() {
33     rand.Seed(time.Now().Unix())
34 
35     fracs := simulate(1_000_000)
36     for i := 2; i <= 12; i++ {
37         fmt.Printf("%2d -> %.2f\n", i, fracs[i])
38     }
39 }
```

Listing 3 shows how we run the simulation. On line 33 we seed the random number generator from the current time to get different results every time we run the program. On line 35 we run the simulation on lines 36-38 we print the results.

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

Listing 4 shows the output of the simulation run. 7 is the winner with 6 & 8 in tie for second place. 


## The Birthday Problem

The [birthday problem](https://en.wikipedia.org/wiki/Birthday_problem) asks what is the probability that in a group of `n` people, at least two people will have the same birthday?

Let's answer this question for a group of 23 people. Try to guess the answer before moving on.

**Listing 5: Checking for same birthday in a random group**

```
10 // hasSame returns True is a random group has the same birthday
11 func hasSame(groupSize int) bool {
12     const daysInYear = 365
13 
14     seen := make(map[int]bool)
15     for i := 0; i < groupSize; i++ {
16         day := rand.Intn(daysInYear)
17         if seen[day] {
18             return true
19         }
20         seen[day] = true
21     }
22 
23     return false
24 }
```

Listing 5 shows the `hasSame` function. For a group of size `n` it will draw `n` random birthdays. Birthdays are represented as days since the beginning of the year, we have 365 of these.

On line 14 we initialized a map for already seen birthdays. On line 16 we draw a random day and on line 17 we check if we've seen this day already. If the currently drawn birthday was already seen, we return `true` on line 18. Otherwise we update the `seen` map on line 20. Finally, if there are no duplicated days, we return `false` on line 23.

Now we can run our simulation.

**Listing 6: Birthday simulation**

```
26 // simulateBirthdays returns the fraction of groups that have at two people
27 // with the same birthday
28 func simulateBirthdays(groupSize, n int) float64 {
29     same := 0
30     for i := 0; i < n; i++ {
31         if hasSame(groupSize) {
32             same++
33         }
34     }
35 
36     return float64(same) / float64(n)
37 }
```

Listing 6 shows the simulation code. We pass the group size and number of iterations as parameters. On line 29 we initialize the number of groups that had at least one duplicate birthday. On line 30 we run check `n` random groups and on line 31 we check if there was a duplicate birthday in the current group. If there is, we update the `same` counter on line 32. Finally on line 36 we return the fraction of groups that had the same birthday.

**Listing 7: Running the simulation**

```
39 func main() {
40     rand.Seed(time.Now().Unix())
41 
42     for i := 0; i < 10; i++ {
43         fmt.Println(simulateBirthdays(23, 1_000_00))
44     }
45 }
```

Listing 7 shows how we run the simulation. On line 40 we seed the random number generator from the current time to get different results every time we run the program. On line 42 we run the simulation 10 times and on line 43 we print the results.

**Listing 8: Output**

```
0.50536
0.50797
0.51021
0.50651
0.50869
0.50819
0.50825
0.50587
0.50783
0.50739
```

Listing 8 shows the output of our 10 simulations. We see that the chance that two people in a group of 23 have the same birthday is about 50% which matches the [expected result](https://en.wikipedia.org/wiki/Birthday_problem).


## Sick or Not?

The following question is taken from Nassim Taleb's [Fooled By Randomness](https://www.amazon.com/Fooled-Randomness-Hidden-Markets-Incerto/dp/0812975219) book which I highly recommend.

> A test of a disease presents a rate of 5% false positives. The disease strikes 1/1000 of the population. People are tested at random, regardless of whether they are suspected of having the disease. A patient’s test is positive. What is the probability of the patient being stricken with the disease?

Try to guess the answer before moving on.

**Listing 9: A random event**

```
10 // probability returns true 1/n of times
11 func probability(n int) bool {
12     return rand.Intn(n) == 1
13 }
```

Listing 9 shows the `probability` function which returns `true` 1/n of the times. On line 12 we compare `rand.Intn` which returns a number between 0 and n, to the number 1. The probability of `rand.Intn(n)` returning 1 (or any other number between 0 and n-1) is 1/n.

**Listing 10: The simulation**

```
15 func simulate(n int) float64 {
16     numSick, numDiagnosed := 0, 0
17     for i := 0; i < n; i++ {
18         sick := probability(1000)
19         if sick {
20             numSick++
21             numDiagnosed++
22         } else {
23             falsePositive := probability(20)
24             if falsePositive {
25                 numDiagnosed++
26             }
27         }
28     }
29 
30     return float64(numSick) / float64(numDiagnosed)
31 }
```

Listing 10 shows the simulation code. On line 16 we initialized the counters for the number of people who are actually sick `numSick` and the number of people who have been diagnosed as sick `numDiagnosed`. On line 18 we draw a random person who might be sick 1/1000 of the times. If the person is sick, we increment both `numSick` and `numDiagnosed` on lines 20 & 21. On line 32, if the person is not sick, we use a 1/20 probability (5%) to see if they were incorrectly diagnosed as sick. If this is the case, we increment `numDiagnosed` on line 25. Finally we return the fraction of people who are actually sick out of the number of people who were diagnosed as sick.

**Listing 11: Running the simulation**

```
33 func main() {
34     rand.Seed(time.Now().Unix())
35     for i := 0; i < 10; i++ {
36         fmt.Println(simulate(1_000_000))
37     }
38 }
```

Listing 11 shows 10 runs of the simulation. One line 34 we initialize the random seed from the current time and on line 35 we run 10 simulations, printing the results on line 36.

**Listing 12: Simulation output**

```
0.020525227299332736
0.020034310729004398
0.019121021356710007
0.02017849453947626
0.01984588399913285
0.019073462324996068
0.01868566904196358
0.019913864633134458
0.020406974486427712
0.018889547520252913
```

Listing 12 shows the output of the 10 simulation runs. The chances that a person is sick given a positive test is about 2% which matches the [expected result](https://psychscenehub.com/psychinsights/well-understand-probabilities-medicine/).

## Conclusion

Simulation is a simple and powerful tool, you don't need to know advanced probability and statistics in order to solve data driven problems. All you need is `math/rand` and some simple logic. After you get your results, you are encouraged to validate them with math. If you don't have the math skills - ask around (I’m asking my friend & college [Shlomo Yona](https://www.mathematic.ai/) whenever I need some math guidance).

If you want to learn more, I recommend watching [Statistics for Hackers](https://www.youtube.com/watch?v=Iq9DzN6mvYA). The examples are in Python, but very easy to follow. You should also read about the [Monte Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method) on Wikipedia and see the wide variety of applications it has.

You can find the code for these simulations [here](https://github.com/353words/simulations).
Catan photo by [Galen Crout](https://unsplash.com/@galen_crout)
