package slim

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/agntcy/slim/control-plane/slimctl/internal/cfg"
	"github.com/agntcy/slim/control-plane/slimctl/internal/cmd"
	"github.com/agntcy/slim/control-plane/slimctl/internal/manager"
)

// NewSlimCmd creates the base 'slim' command.
func NewSlimCmd(ctx context.Context, appConfig *cfg.AppConfig) *cobra.Command {
	slimCmd := &cobra.Command{
		Use:     "slim",
		Aliases: []string{"s"},
		Short:   "Manage a local SLIM instance",
		Long:    `Manage a local SLIM instance`,
		GroupID: cmd.GroupSlim,
	}
	// Pass appConfig to subcommands so that they can create managers with proper configuration.
	slimCmd.AddCommand(NewStartCmd(appConfig))
	// slimCmd.AddCommand(NewStopCmd(appConfig))
	// slimCmd.AddCommand(NewStatusCmd(appConfig))
	return slimCmd
}

// NewStartCmd creates the 'start' command.
func NewStartCmd(appConfig *cfg.AppConfig) *cobra.Command {
	var endpoint string
	var port string

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the SLIM instance",
		Long:  `Start the SLIM instance`,
		RunE: func(c *cobra.Command, args []string) error {
			// If no endpoint specified, bind to loopback on the given port
			if endpoint == "" {
				endpoint = "127.0.0.1:" + port
			}
			mgr := manager.NewManager(appConfig.CommonOpts.Logger, endpoint, port)
			return mgr.Start(c.Context())
		},
	}

	startCmd.Flags().StringVar(&endpoint, "endpoint", "", "Endpoint to bind (default: 127.0.0.1:<port>)")
	startCmd.Flags().StringVar(&port, "port", "8080", "Port to listen on")

	return startCmd
}

// NewStopCmd creates the 'stop' command.
func NewStopCmd(appConfig *cfg.AppConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stop the SLIM instance",
		Long:  `Stop the SLIM instance`,
		RunE: func(c *cobra.Command, args []string) error {
			mgr := manager.NewManager(appConfig.CommonOpts.Logger, "", "")
			return mgr.Stop(c.Context())
		},
	}
}

// NewStatusCmd creates the 'status' command.
func NewStatusCmd(appConfig *cfg.AppConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get the status of the SLIM instance",
		Long:  `Get the status of the SLIM instance`,
		RunE: func(c *cobra.Command, args []string) error {
			mgr := manager.NewManager(appConfig.CommonOpts.Logger, "", "")
			status, err := mgr.Status(c.Context())
			if err != nil {
				return fmt.Errorf("failed to get status: %w", err)
			}
			fmt.Printf("SLIM instance status: %s\n", status)
			return nil
		},
	}
}
