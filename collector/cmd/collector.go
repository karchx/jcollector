package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var collectorCmd = &cobra.Command{
	Use:   "collector",
	Short: "Run collector file system linux (/var/log)",
	Run:   runCollector,
}

// LogEntry line log normalize
type LogEntry struct {
	Timestamp   time.Time              `json:"timestmap"`
	Source      string                 `json:"source"`
	Level       string                 `json:"level"`
	Message     string                 `json:"message"`
	Metadata    map[string]interface{} `json:"metadata"`
	CollectedAt time.Time              `json:"collector_at"`
}

func runCollector(cmd *cobra.Command, args []string) {
	fmt.Println("Hola mundo!")
}

func init() {
	rootCmd.AddCommand(collectorCmd)
}
