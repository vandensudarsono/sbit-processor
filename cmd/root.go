package cmd

import (
	"fmt"
	"sbit-processor/cmd/action"
	"sbit-processor/config"
	logger "sbit-processor/infrastructure/log"

	"github.com/spf13/cobra"
)

// CommandEngine is the structure of cli
type Command struct {
	rootCmd *cobra.Command
}

var text = "SBIT-PROCESSOR"

// NewCommandEngine the command line boot loader
func NewCommand() *Command {
	var rootCmd = &cobra.Command{
		Use:   "sbit-processor",
		Short: "sbit processor service command line",
		Long:  "sbit processor service command line",
	}

	return &Command{
		rootCmd: rootCmd,
	}
}

// Run the all command line
func (c *Command) Run() {
	//container := containers.NewEngine()
	var rootCommands = []*cobra.Command{
		{
			Use:   "serve",
			Short: "Run sbit-processor service",
			Long:  "Run sbit-processor HTTP service",
			PreRun: func(cmd *cobra.Command, args []string) {
				// initialize config
				config.LoadConfig()

				// Show display text
				fmt.Println(text)

				logger.WithFields(logger.Fields{"component": "command", "action": "serve with watcher"}).
					Infof("PreRun command done")
			},
			Run: func(cmd *cobra.Command, args []string) {
				action.RunWalletProcessor()
			},
			PostRun: func(cmd *cobra.Command, args []string) {
				logger.WithFields(logger.Fields{"component": "command", "action": "serve with watcher"}).
					Infof("PostRun command done")
			},
		},
	}

	for _, command := range rootCommands {
		c.rootCmd.AddCommand(command)
	}

	c.rootCmd.Execute()
}

// GetRoot the command line service
func (c *Command) GetRoot() *cobra.Command {
	return c.rootCmd
}
