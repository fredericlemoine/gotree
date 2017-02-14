package cmd

import (
	"github.com/spf13/cobra"
)

// clearsupportCmd represents the clearsupport command
var clearpvalueCmd = &cobra.Command{
	Use:   "pvalues",
	Short: "Clear pvalues associated to supports from input trees",
	Long:  `Clear pvalues associated to supports from input trees.`,
	Run: func(cmd *cobra.Command, args []string) {
		f := openWriteFile(outtreefile)
		for t := range readTrees(intreefile) {
			t.Tree.ClearPvalues()
			f.WriteString(t.Tree.Newick() + "\n")
		}
		f.Close()
	},
}

func init() {
	clearCmd.AddCommand(clearpvalueCmd)
}