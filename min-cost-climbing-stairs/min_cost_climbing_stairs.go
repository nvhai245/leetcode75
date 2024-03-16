package mincostclimbingstairs

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	minCost := make([]int, len(cost)+1)
	minCost[len(cost)], minCost[len(cost)-1] = 0, cost[len(cost)-1]
	for i := len(cost) - 2; i >= 0; i-- {
		minCost[i] = cost[i] + min(minCost[i+1], minCost[i+2])
	}
	return min(minCost[0], minCost[1])
}
