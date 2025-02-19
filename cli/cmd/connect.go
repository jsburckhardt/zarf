package cmd

import (
	"github.com/defenseunicorns/zarf/cli/internal/k8s"
	"github.com/spf13/cobra"
)

var (
	connectResourceName string
	connectNamespace    string
	connectResourceType string
	connectLocalPort    int
	connectRemotePort   int
	cliOnly   bool

	connectCmd = &cobra.Command{
		Use:     "connect <REGISTRY|LOGGING|GIT>",
		Aliases: []string{"c"},
		Short:   "Access services or pods deployed in the cluster.",
		Run: func(cmd *cobra.Command, args []string) {
			var target string
			if len(args) > 0 {
				target = args[0]
			}
			tunnel := k8s.NewTunnel(connectNamespace, connectResourceType, connectResourceName, connectLocalPort, connectRemotePort)
			// If the cliOnly flag is false (default), enable auto-open
			if !cliOnly {
				tunnel.EnableAutoOpen()
			}
			tunnel.Connect(target, true)
		},
	}
)

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringVar(&connectResourceName, "name", "docker-registry", "Specify the resource name.  E.g. name=unicorns or name=unicorn-pod-7448499f4d-b5bk6")
	connectCmd.Flags().StringVar(&connectNamespace, "namespace", k8s.ZarfNamespace, "Specify the namespace.  E.g. namespace=default")
	connectCmd.Flags().StringVar(&connectResourceType, "type", k8s.SvcResource, "Specify the resource type.  E.g. type=svc or type=pod")
	connectCmd.Flags().IntVar(&connectLocalPort, "local-port", 0, "(Optional, autogenerated if not provided) Specify the local port to bind to.  E.g. local-port=42000")
	connectCmd.Flags().IntVar(&connectRemotePort, "remote-port", 0, "Specify the remote port of the resource to bind to.  E.g. remote-port=8080")
	connectCmd.Flags().BoolVar(&cliOnly, "cli-only", false, "Disable browser auto-open")
}
