package main

import (
	"fmt"
	"github.com/phob0s-pl/weatherllo/owm"
	"github.com/spf13/cobra"
)

func version(cmd *cobra.Command, args []string) {
	fmt.Println(owm.Version)
}
