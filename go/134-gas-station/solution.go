/*
The most important detail is that there's always one solution
This means we don't have to go through the whole cycle:
If we find an index start that allows to go all the way to the end of there
gas array: we can assume that the rest of the trip is valid
	1. Store the amount of gas in `tank`
	2. Check that sum(gas) >= sum(cost), else: return -1
		a. If sum(cost) <= sum(gas): there's a unique solution
	3. Set the default result to 0
	4. Iterate over all elements in gas:
		a. Update tank value: tank = gas[i] - cost[i]
		b. If tank > 0
			- we keep on this track
		c. else
			- reset tank
			- set start to next item: start = start + i

*/
package main

func sum(slice []int) int {
	var sum = 0
	for _, i := range slice {
		sum += i
	}

	return sum
}

func canCompleteCircuit(gas []int, cost []int) int {
	var tank = 0
	if sum(gas) < sum(cost) {
		return -1
	}

	var result = 0
	for i := 0; i < len(gas); i++ {
		tank = tank + gas[i] - cost[i]
		if tank < 0 {
			tank = 0
			result = i + 1
		}
	}

	return result
}
