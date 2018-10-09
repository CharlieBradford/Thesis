
package main

import "fmt"
import "math/rand"
import "math"

type V struct
{
	index int
	trust float64
	votes []int
}

var baseVoters int = 3 // baseVoters real voters and 
var fakeMulti int = 3 // fakeMulti * baseVoters fake voters
var alpha float64 = 3

var nElections int = 7

var voterPattern = [][]int{
	{1, 6, 5, 8, 4, 2, 1},
	{1, 1, 5, 9, 1, 4, 1},
	{2, 6, 4, 8, 2, 4, 1},
	{1, 3, 4, 8, 4, 4, 1},
	{2, 6, 5, 1, 3, 2, 1}}

func main() {
	votes := generateVotes()
	reciprocal(votes)

}

func affine(voters []V) {



}

func reciprocal(voters []V) {
	nVoters := 5 * baseVoters * (fakeMulti + 1)
	est := initialEstimate(voters)
	fmt.Println(est)
	i := 0
	diff := 1.0

	var rank [][]float64
	var normRank [][]float64

	for{
		fmt.Println("***********************************************")
		old := diff
		rank = [][]float64{
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}}
		normRank = [][]float64{
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}}

		j := 0
		for{
			k := 0
			for{
				rank[k][voters[j].votes[k]-1] += math.Pow(voters[j].trust,alpha)
				k += 1
				if k >= 7 { break }
			}
			j += 1
			if j >= nVoters { break }
		}

		j = 0
		for{
			sum := 0.0
			k := 0
			for{
				sum += math.Pow(rank[j][k], 2)
				k += 1
				if k >= 9 { break }
			}
			norm := math.Pow(sum, 0.5)
			k = 0
			for{
				normRank[j][k] = rank[j][k]/norm
				k += 1
				if k >= 9 { break }
			}
			fmt.Println(normRank[j])
			j += 1
			if j >= nElections { break }
		}

		j = 0
		for{
			voters[j].trust = 0.0
			k := 0
			for{
				voters[j].trust += normRank[k][voters[j].votes[k]-1]
				k += 1
				if k >= nElections { break }
			}
			fmt.Println("Iteration: ",i," Voter: ",j," Trust: ",voters[j].trust)
			j += 1
			if j >= nVoters { break }
		}

		diff = 0
		j = 0
		for{
			k := 1
			for{
				diff += normRank[j][k]
				k += 1
				if k >= 9 { break }
			}
			j += 1
			if j >= nElections { break }
		}

		i += 1

		if math.Pow(old - diff, 0.5) < 0.0001 { break }
	}

	i = 0
	for{
		max := 0.0
		index := -1
		j := 0
		for{
			if normRank[i][j] > max {
				index = j
				max = normRank[i][j]
			}
			j += 1
			if j >= 9 { break }
		}
		fmt.Println("Election", i+1, "winner is", index+1)
		i += 1
		if i >= nElections { break }
	}





}

func initialEstimate(voters []V) []float64 {
	nVoters := 5 * baseVoters * (fakeMulti + 1)
	i := 0

	var est []float64
	for {
		j := 0
		sum := 0.0
		for{
			sum += float64(voters[j].votes[i])
			j += 1
			if j >= nVoters { break }
		}
		est = append(est, sum/float64(nVoters))
		i += 1
		if i >= nElections { break }
	}
	return est

}

func generateVotes() []V {
	var voters []V
	i := 0
	for {
		if i < baseVoters * 5 {
			voter := V{i, 1.0, voterPattern[i%5]}
			voters = append(voters, voter)
		} else if i < baseVoters * (1 + fakeMulti) * 5 {
			pattern := []int{rand.Intn(9)+1, rand.Intn(9)+1, rand.Intn(9)+1,
					 rand.Intn(9)+1, rand.Intn(9)+1, rand.Intn(9)+1,
					 5}
			voter := V{i, 1.0, pattern}
			voters = append(voters, voter)
		} else {
			break;
		}
		i += 1
	}
	return voters
}

