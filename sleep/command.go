package sleep

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var (
	DelayMs int64
)

var Command = &cobra.Command{
	Use:   "sleep",
	Short: "Sleep command",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(fmt.Sprintf("Before sleeping for %dms", DelayMs))
		time.Sleep(time.Duration(DelayMs) * time.Millisecond)
		log.Println("done")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
	Command.Flags().Int64Var(&DelayMs, "delayms", 10000, "delayms")
}
