package main

import "github.com/spf13/cobra"

func createCommandsStructure() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "weatherllo",
		Short: "weatherllo is a openweathermap.org client",
		Long:  `weatherllo is a openweathermap.org client`,
		Run:   nil,
	}

	rootCmd.AddCommand(createListenCommand())
	rootCmd.AddCommand(createVersionCommand())
	return rootCmd
}

func createListenCommand() *cobra.Command {
	listenCmd := &cobra.Command{
		Use:   "listen",
		Short: "start listener",
		Long:  `start listener`,
		Run:   listen,
	}

	listenCmd.Flags().Bool(debugFlag, false, "debug logs")
	listenCmd.Flags().String(addrFlag, ":8080", "listen address")
	listenCmd.Flags().String(tokenFlag, "", "openweathermap.org API token")

	return listenCmd
}

func createVersionCommand() *cobra.Command {
	listenCmd := &cobra.Command{
		Use:   "version",
		Short: "print application version",
		Long:  `print application version`,
		Run:   version,
	}

	return listenCmd
}
