package day9

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
		sum := 0
		for scanner.Scan() {
			seq, err := ParseSequence(scanner.Text())
			if err != nil {
				return err
			}

			sum += NextVal(seq)
		}

		fmt.Println(sum)

		return nil
	},
}

func NextVal(seq []int) int {
	diffs := Differentiate(seq)
	if AllZero(diffs) {
		return seq[0]
	}

	return seq[len(seq)-1] + NextVal(diffs)
}
