package day8

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var task1Cmd = &cobra.Command{
	Use:   "task1 <inputFile>",
	Short: "Solves task 1 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(args[0])
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		directions := ParseDirections(scanner)
		nodes := ParseNodes(scanner)

		steps, err := WalkPath(directions, nodes["AAA"], func(node *Node) bool {
			return node.Name == "ZZZ"
		})
		if err != nil {
			return err
		}

		fmt.Println(steps)

		return nil
	},
}
