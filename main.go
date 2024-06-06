package main

import (
	"leetcode-go/dweek/dweek130"
	"leetcode-go/dweek/dweek131"
	"leetcode-go/easy"
	"leetcode-go/hard"
	"leetcode-go/lcp"
	"leetcode-go/lcr"
	"leetcode-go/middle"
	"leetcode-go/week/week336"
	"leetcode-go/week/week392"
	"leetcode-go/week/week393"
	"leetcode-go/week/week394"
	"leetcode-go/week/week395"
	"leetcode-go/week/week398"
	"leetcode-go/week/week399"
)

func main() {

	week336.BeautifulSubarrays([]int{4, 3, 1, 2, 4})

	easy.MinNumberOfHours(5, 1, []int{1, 3, 3}, []int{1, 3, 7})

	middle.MaximalNetworkRank(4,
		[][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}})

	easy.ArithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3)

	middle.IsRobotBounded("GGLLGG")

	middle.NumberOfWeakCharacters([][]int{{5, 5}, {6, 3}, {3, 6}, {4, 4}})

	easy.MostFrequentEven([]int{0, 1, 2, 2, 4, 4, 1})
	easy.SortPeople([]string{"IEO", "Sgizfdfrims", "QTASHKQ", "Vk", "RPJOFYZUBFSIYp", "EPCFFt", "VOYGWWNCf", "WSpmqvb"},
		[]int{17233, 32521, 14087, 42738, 46669, 65662, 43204, 8224})

	middle.ReverseList(&middle.ListNode{Val: 1, Next: &middle.ListNode{Val: 2, Next: &middle.ListNode{Val: 3, Next: &middle.ListNode{Val: 4, Next: nil}}}})

	middle.HalveArray([]int{5, 19, 8, 1})

	middle.RepairCars([]int{4, 2, 3, 1}, 10)

	middle.DistanceK(&middle.TreeNode{Val: 1, Left: &middle.TreeNode{Val: 2, Left: &middle.TreeNode{Val: 4, Left: nil, Right: nil}, Right: nil}, Right: &middle.TreeNode{Val: 3, Left: nil, Right: &middle.TreeNode{Val: 5, Left: nil, Right: nil}}}, &middle.TreeNode{Val: 1, Left: nil, Right: nil}, 1)

	middle.KthLargestNumber([]string{"1", "2", "0", "7", "0", "2", "0"}, 4)

	middle.QueensAttacktheKing([][]int{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}}, []int{3, 3})

	easy.CountPoints("B0B6G0R6R0R6G9")

	middle.MinPathCost([][]int{{5, 3}, {4, 0}, {2, 1}}, [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}})

	//[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
	middle.BstToGst(&middle.TreeNode{
		Val: 4,
		Left: &middle.TreeNode{
			Val: 1,
			Left: &middle.TreeNode{
				Val: 0,
			},
			Right: &middle.TreeNode{
				Val: 2,
				Right: &middle.TreeNode{
					Val: 3,
				},
			}},
		Right: &middle.TreeNode{
			Val: 6,
			Left: &middle.TreeNode{
				Val: 5,
			},
			Right: &middle.TreeNode{
				Val: 7,
				Right: &middle.TreeNode{
					Val: 8,
				},
			}},
	})

	middle.ReverseOddLevels(
		&middle.TreeNode{
			Val: 0,
			Left: &middle.TreeNode{
				Val: 1,
				Left: &middle.TreeNode{
					Val: 0,
					Left: &middle.TreeNode{
						Val: 1,
					},
					Right: &middle.TreeNode{
						Val: 1,
					},
				},
				Right: &middle.TreeNode{
					Val: 0,
					Left: &middle.TreeNode{
						Val: 1,
					},
					Right: &middle.TreeNode{
						Val: 1,
					},
				},
			},
			Right: &middle.TreeNode{
				Val: 2,
				Left: &middle.TreeNode{
					Val: 0,
					Left: &middle.TreeNode{
						Val: 2,
					},
					Right: &middle.TreeNode{
						Val: 2,
					},
				},
				Right: &middle.TreeNode{
					Val: 0,
					Left: &middle.TreeNode{
						Val: 2,
					},
					Right: &middle.TreeNode{
						Val: 2,
					},
				},
			},
		},
	)

	middle.MaximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2)

	middle.NumOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 10)

	lcp.MagicTower([]int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150})

	lcr.LenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8})

	middle.CountWays([][]int{{1, 3}, {10, 20}, {2, 5}, {4, 8}})

	easy.MinimumSum([]int{8, 6, 1, 5, 3})

	week392.GetSmallestString("xaxcd", 4)
	week392.MinOperationsToMakeMedianK([]int{1, 2, 3, 4, 5, 6}, 4)

	middle.MaximumBinaryString("01111001100000110010")

	hard.MinOperations([]int{4, 2, 3, 5})

	middle.Divide2(-1021989372, -82778243)
	week393.FindLatestTime("?0:40")
	middle.CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	week394.MinimumOperations([][]int{{0, 6, 2}, {9, 0, 9}, {4, 9, 6}})

	middle.MaxSatisfied([]int{10, 1, 7}, []int{0, 0, 0}, 2)

	lcr.AddBinary("11", "10")

	lcr.SingleNumber([]int{-2})

	cache := lcr.ConstructorLRU(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
	cache.Get(2)
	cache.Put(4, 4)
	cache.Get(1)
	cache.Get(3)
	cache.Get(4)

	snapArray := middle.Constructor1146(1)
	snapArray.Snap()
	snapArray.Snap()
	snapArray.Set(0, 4)
	snapArray.Snap()
	snapArray.Get(0, 1)
	snapArray.Set(0, 12)
	snapArray.Get(0, 1)
	snapArray.Snap()
	snapArray.Get(0, 3)

	week395.MinimumAddedInteger([]int{9, 4, 3, 9, 4}, []int{7, 8, 8})

	middle.TotalCost([]int{2, 2, 2, 2, 2, 2, 1, 4, 5, 5, 5, 5, 5, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 7, 3)

	lcr.NumSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)

	lcr.LengthOfLongestSubstring("abcabcbb")

	lcr.MinWindow("ADOBECODEBANC", "ABC")

	lcr.IsPalindrome("A man, a plan, a canal: Panama")

	middle.GarbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3})

	dweek130.MinimumSubstringsInPartitionII("fabccddg")

	lcr.Partition("google")
	lcr.GenerateParenthesis(3)

	middle.NumberOfWeeks([]int{5, 7, 5, 7, 9, 7})
	lcr.FindMinDifference([]string{"00:00", "23:59", "00:00"})

	lcr.FindCircleNum([][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}})

	lcr.PathTarget(
		&lcr.TreeNode{
			Val: 5,
			Left: &lcr.TreeNode{
				Val: 4,
				Left: &lcr.TreeNode{
					Val:   11,
					Left:  &lcr.TreeNode{Val: 7},
					Right: &lcr.TreeNode{Val: 2},
				}},
			Right: &lcr.TreeNode{
				Val:  8,
				Left: &lcr.TreeNode{Val: 13},
				Right: &lcr.TreeNode{
					Val:   4,
					Left:  &lcr.TreeNode{Val: 5},
					Right: &lcr.TreeNode{Val: 1},
				}},
		},
		22)

	week398.SumDigitDifferences([]int{13, 23, 12})

	middle.MostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4)

	dweek131.QueryResults(4, [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}})

	week399.NumberOfPairs([]int{1, 2}, []int{3,2}, 2)

	middle.MaximumLength("abcdef")

	hard.CherryPickupII([][]int{{0,1,-1},{1,0,-1},{1,1,1}})

	easy.DistributeCandies(10,3)

}
