package main

import (
	"log"

	"github.com/ryrden/birdie-challenge/internal/app/handler"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "transform [flags] - birdie-challenge",
	Short: "Transform operations on a JSON file.",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := handler.TransformJSON(args); err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.birdie-challenge.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func main() {
	Execute()
}



