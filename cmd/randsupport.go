package cmd

import (
	"github.com/fredericlemoine/gostats"
	"github.com/spf13/cobra"
	"math/rand"
	"time"
)

// randsupportCmd represents the randbrlen command
var randsupportCmd = &cobra.Command{
	Use:   "randsupport",
	Short: "Assign a random support to edges of input trees",
	Long: `Assign a random support to edges of input trees.

Support follows a uniform distribution in [0,1].

`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		rand.Seed(seed)
	},
	Run: func(cmd *cobra.Command, args []string) {
		f := openWriteFile(outtreefile)
		for tr := range readTrees(intreefile) {
			for _, e := range tr.Tree.Edges() {
				if !e.Right().Tip() {
					e.SetSupport(gostats.Float64Range(0, 1))
				}
			}
			f.WriteString(tr.Tree.Newick() + "\n")
		}
		f.Close()
	},
}

func init() {
	RootCmd.AddCommand(randsupportCmd)

	randsupportCmd.PersistentFlags().StringVarP(&intreefile, "input", "i", "stdin", "Input tree")
	randsupportCmd.PersistentFlags().StringVarP(&outtreefile, "output", "o", "stdout", "Output file")
	randsupportCmd.Flags().Int64VarP(&seed, "seed", "s", time.Now().UTC().UnixNano(), "Initial Random Seed")
}