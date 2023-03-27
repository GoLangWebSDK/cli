package cli

import (
	"os"

	"github.com/spf13/cobra"
)

type CLICmds interface {
	AddServiceCommand(serviceCmd *cobra.Command)
}

type CLIServer interface {
	Execute()
}

type CLI struct {
	ServiceCommand *cobra.Command
	Server         CLIServer
}

func NewCLI(config any) *CLI {
	cli := &CLI{
		ServiceCommand: &cobra.Command{
			Use:   "service",
			Short: "Command line manager for service.",
			Long:  `Command line manager for service modules`,
		},
	}

	cli.ServiceCommand.Run = cli.runServiceCommand
	return cli
}

func (cli *CLI) runServiceCommand(_ *cobra.Command, _ []string) {
	cli.Server.Execute()
}

func (cli *CLI) AddCommand(cmd *cobra.Command) {
	cli.ServiceCommand.AddCommand(cmd)
}

func (cli *CLI) AddCommands(CLICmds ...CLICmds) {
	for _, cliCmd := range CLICmds {
		cliCmd.AddServiceCommand(cli.ServiceCommand)
	}
}

func (cli *CLI) LoadCommands(c CLICmds) {
	c.AddServiceCommand(cli.ServiceCommand)
}

func (cli *CLI) Execute(server CLIServer) {
	cli.Server = server
	if err := cli.ServiceCommand.Execute(); err != nil {
		// config.AppSetting.Logger.Fatal().Err(err).Msg("Failed to execute main service command")
		os.Exit(1)
	}
}
