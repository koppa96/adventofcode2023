package day8

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var task2Cmd = &cobra.Command{
	Use:   "task2 <inputFile>",
	Short: "Solves task 2 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		src := []int{1, 2, 3}
		var dst []int

		copy(dst, src)

		file, err := os.Open(args[0])
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		directions := ParseDirections(scanner)
		nodes := ParseNodes(scanner)

		var startingNodes []*Node
		for _, node := range nodes {
			if node.Name[len(node.Name)-1] == 'A' {
				startingNodes = append(startingNodes, node)
			}
		}

		times, err := getPathTimes(directions, startingNodes)
		if err != nil {
			return err
		}

		lcm := LCM(times)

		fmt.Println(lcm * len(directions))

		return nil
	},
}

// See https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers []int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(append([]int{result}, integers[i:]...))
	}

	return result
}

func getPathTimes(directions []rune, startingNodes []*Node) ([]int, error) {
	results := make(chan int)
	errors := make(chan error)

	for i := range startingNodes {
		node := startingNodes[i]

		go func() {
			length, err := WalkPath(directions, node, func(node *Node) bool {
				return node.Name[len(node.Name)-1] == 'Z'
			})

			if err != nil {
				errors <- err
			}

			results <- length
		}()
	}

	var lengths []int
	for len(lengths) < len(startingNodes) {
		select {
		case err := <-errors:
			return nil, err
		case length := <-results:
			lengths = append(lengths, length/len(directions))
		}
	}

	return lengths, nil
}

func allEqual(items []int) bool {
	if len(items) == 0 {
		return true
	}

	first := items[0]
	for i := 1; i < len(items); i++ {
		if items[i] != first {
			return false
		}
	}

	return true
}
