package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*

--- Day 1: Chronal Calibration ---
"We've detected some temporal anomalies," one of Santa's Elves at the Temporal Anomaly Research and Detection Instrument Station tells you. She sounded pretty worried when she called you down here. "At 500-year intervals into the past, someone has been changing Santa's history!"

"The good news is that the changes won't propagate to our time stream for another 25 days, and we have a device" - she attaches something to your wrist - "that will let you fix the changes with no such propagation delay. It's configured to send you 500 years further into the past every few days; that was the best we could do on such short notice."

"The bad news is that we are detecting roughly fifty anomalies throughout time; the device will indicate fixed anomalies with stars. The other bad news is that we only have one device and you're the best person for the job! Good lu--" She taps a button on the device and you suddenly feel like you're falling. To save Christmas, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

After feeling like you've been falling for a few minutes, you look at the device's tiny screen. "Error: Device must be calibrated before first use. Frequency drift detected. Cannot maintain destination lock." Below the message, the device shows a sequence of changes in frequency (your puzzle input). A value like +6 means the current frequency increases by 6; a value like -3 means the current frequency decreases by 3.

For example, if the device displays frequency changes of +1, -2, +3, +1, then starting from a frequency of zero, the following changes would occur:

Current frequency  0, change of +1; resulting frequency  1.
Current frequency  1, change of -2; resulting frequency -1.
Current frequency -1, change of +3; resulting frequency  2.
Current frequency  2, change of +1; resulting frequency  3.
In this example, the resulting frequency is 3.

Here are other example situations:

+1, +1, +1 results in  3
+1, +1, -2 results in  0
-1, -2, -3 results in -6
Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?

Your puzzle answer was 442.

--- Part Two ---
You notice that the device repeats the same frequency change list over and over. To calibrate the device, you need to find the first frequency it reaches twice.

For example, using the same list of changes above, the device would loop as follows:

Current frequency  0, change of +1; resulting frequency  1.
Current frequency  1, change of -2; resulting frequency -1.
Current frequency -1, change of +3; resulting frequency  2.
Current frequency  2, change of +1; resulting frequency  3.
(At this point, the device continues from the start of the list.)
Current frequency  3, change of +1; resulting frequency  4.
Current frequency  4, change of -2; resulting frequency  2, which has already been seen.
In this example, the first frequency reached twice is 2. Note that your device might need to repeat its list of frequency changes many times before a duplicate frequency is found, and that duplicates might be found while in the middle of processing the list.

Here are other examples:

+1, -1 first reaches 0 twice.
+3, +3, +4, -2, -4 first reaches 10 twice.
-6, +3, +8, +5, -6 first reaches 5 twice.
+7, +7, -2, -7, -4 first reaches 14 twice.
What is the first frequency your device reaches twice?

*/

func main() {
	fmt.Println("Day 1 - Chronal Calibration")
	sum, freq := ChronalCalibration(puzzleInput)
	fmt.Println(" > Sum:", sum)
	fmt.Println(" > Freq:", freq)
}

func ChronalCalibration(in string) (int, int) {
	sum := 0
	list := []int{}
	for _, l := range strings.Split(in, "\n") {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		list = append(list, i)
		sum += i
	}

	repeatCount := map[int]int{0: 1}
	i := 0
	freq := 0
	for {
		freq += list[i]
		repeatCount[freq]++
		if repeatCount[freq] > 1 {
			return sum, freq
		}
		i++
		if i >= len(list) {
			i = 0
		}
	}
}

const puzzleInput = `-4
+7
+3
+1
-9
-14
+18
-7
-5
-18
+11
-8
+17
-16
-19
+14
+11
-8
+14
+22
+13
+14
-18
+8
-16
+10
-12
+9
-19
-12
-6
+10
+2
-14
+18
+17
+11
-5
+6
+9
+16
-3
+12
+5
+15
+7
+2
-5
-13
+7
+19
+10
-2
+3
-5
-7
-11
+14
+13
+3
+11
+15
-19
-1
-5
+15
+14
-16
+8
+9
-11
-7
+15
+4
+7
+11
-2
+17
-8
-14
+3
-4
+18
+1
+6
-5
+17
+13
-14
-15
-9
+16
+14
+12
-14
-14
+4
-19
+11
+5
+7
+1
+13
-7
+19
+12
+10
+13
-3
+6
-17
+13
-8
+16
+4
-15
+14
-1
+15
-19
+15
-19
+17
-6
+11
-10
+9
+17
+6
+15
-18
+2
-8
+11
-6
-7
+9
+16
-13
-18
-1
+5
-16
+9
-6
+11
-16
+3
-9
+18
+5
+7
+15
+12
+19
-17
-4
-15
-13
+5
-11
+25
+14
+5
+5
+4
-2
-4
-14
+7
+16
+6
+11
-8
+13
+8
+3
-12
+8
+7
-5
+14
+5
+8
-3
-14
-14
-7
-18
-12
+18
+4
-14
-17
-17
+12
-9
+3
-18
+27
+26
-9
-7
+12
+9
+6
+7
+1
-13
+4
-8
-3
-12
+21
+13
+6
+7
+9
+18
-17
+4
+10
-11
-17
-16
+18
+9
-2
+3
-13
-9
+3
-10
-15
+11
+24
+2
+16
+11
+18
+14
-5
-3
+18
+4
+4
+5
-2
-10
+8
+19
+4
+14
+8
+1
-14
-7
+15
-4
-5
+20
+15
+2
-14
-12
-16
+4
-9
+18
+16
+7
+7
+10
+17
-18
+3
+10
-11
-21
-9
-2
+12
-17
+13
+19
-18
-4
-19
+10
-11
-18
-4
+6
+1
+2
-4
-28
-28
-19
+13
+18
-15
-6
+5
+9
-26
-6
+19
-6
+1
-11
-11
+25
-26
+19
+10
+2
+9
-22
-66
-8
+1
-14
-2
-5
+19
-55
+6
-19
-4
-13
+21
-13
-3
-10
-7
+15
-5
-4
-10
+5
-11
-3
+16
+6
+10
+3
-1
-5
-21
-6
+9
-4
-6
-17
-10
+14
-11
+17
+16
-19
+6
+1
-2
+8
-15
+17
-14
-16
+7
-5
-24
-7
-20
+5
-4
+11
-1
+12
-9
-11
-9
-9
-11
-3
-15
-21
-12
+21
+3
+12
+2
+2
+21
-9
+20
+4
+11
-1
+4
-20
+4
+19
+12
+4
-15
+14
+11
+21
-29
-53
+10
-11
+8
-10
-11
+12
-17
-25
-35
+1
-2
+14
+36
+70
+14
+50
+3
+6
+23
-20
-6
-66
-39
-14
+191
-17
+24
+155
-6
+61
+59623
-4
+7
+19
+5
-1
-8
-8
-15
+16
+8
-7
+19
-3
+14
+14
+6
+17
-15
+3
+11
+12
-5
+19
+2
+1
+16
+13
+6
+1
+18
+16
-4
+11
-1
+19
-7
-19
-10
+19
+7
+4
-17
-4
-4
+9
-8
-17
-9
-14
+16
+4
-19
+11
+1
+2
-19
+7
-8
-7
-18
-6
+4
-5
-15
+2
+7
-18
-8
+14
-20
+1
+17
-2
+3
+17
-1
+7
+10
+10
-1
+19
+3
-16
-10
+1
+10
+6
+20
-17
+19
+14
-18
-11
+22
-3
-16
+15
+12
+14
-12
-1
+22
+10
-7
-8
+2
-1
-6
-7
+19
+7
+12
-14
+18
+7
-17
-9
+10
+14
-1
-17
+8
+4
+2
+16
+3
+16
+4
-2
+18
-19
-15
-1
-10
-5
+8
-4
-13
+18
+15
-5
+10
-6
+13
-11
+12
+2
+4
-2
-13
-7
-11
-13
-5
+4
+10
+12
-15
-16
-10
-9
-27
-4
-9
-21
-13
-4
-1
-10
-10
+14
+12
-29
-16
-12
-1
-12
+18
-20
-11
-5
+12
-2
-6
+18
-2
-1
+6
-18
-12
+20
-4
+6
-14
+6
-17
-6
+8
-14
-18
-4
-8
+2
+7
+20
-10
-15
+18
-12
+13
-18
+9
+11
-19
+9
+3
+10
-17
-16
-4
-11
-5
-9
-12
+4
-13
+14
+13
-17
-21
+7
+19
+10
-6
+9
+13
-5
-13
+19
+7
-4
-6
+17
-16
-8
-14
-11
-3
-16
+10
-1
-2
-19
+3
+13
-15
+8
+6
-19
+2
-12
-6
-19
+17
-8
+11
+2
-18
-6
-10
+7
+4
+12
-3
+8
+7
+18
-8
+10
-15
-2
-12
+20
-13
-16
-11
+9
-1
+11
+19
-5
+20
-2
+17
+17
+17
+18
+17
+25
-31
+20
-7
+21
+15
+15
-13
+18
+14
+28
-10
+16
-17
-16
-2
-7
-8
-23
-16
+11
-20
+11
-3
-31
-4
-20
+18
-8
+3
+6
-15
-35
+19
+2
-4
+9
-12
-15
-9
-3
+15
+38
-6
+11
-36
-11
-14
-1
-5
+2
+23
-12
-33
-18
+1
+5
+3
+16
+7
-28
+8
-22
-7
-18
-17
+16
+11
+16
+4
-15
+7
-10
-7
-9
-4
+9
+6
-21
-11
-20
-19
-5
+3
+27
+9
+23
+16
-19
-2
+20
+13
-2
-16
+11
+2
-10
-11
-21
-6
-8
-10
-20
+35
+46
+2
-19
-13
+23
-2
+56
+7
+53
+5
-64
-62
+103
+162
+4
+10
-27
-15
-14
+22
+9
+1
+38
-9
+8
+9
+34
+1
-11
+4
-19
-29
+16
+52
-16
+27
+188
-83
+191
+59071
+19
-10
-6
-16
-18
+14
-10
+18
+4
-10
+3
+16
+6
-18
-3
+1
+4
-20
+4
-16
-7
-6
+19
+12
-3
-7
-13
-7
-9
+19
-20
+12
+12
+18
-15
-8
-11
-9
-7
+13
+16
+11
+4
-13
-10
+1
-11
-14
+13
+4
+10
-5
+7
-24
-7
+11
-22
-20
+15
-10
+13
-7
-25
-2
-4
+45
-1
+16
+11
+24
+4
+4
+22
-17
+13
-16
+5
+16
+16
+15
-119289`
