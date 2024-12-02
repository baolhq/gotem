package lib

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ImportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "import",
		Short: "Import files and directories into stash",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Imporint..")
		},
	}
}

func TestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Test program",
		Run: func(cmd *cobra.Command, args []string) {
      Decode()
		},
	}
}
