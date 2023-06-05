package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   Name,
		Short: "A kratos-layout service",
		Long: `A kratos-layout service. 
				Complete documentation is available at https://omalloc.com/kratos-layout`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  `All software has versions. My has one too.`,
		Run: func(cmd *cobra.Command, args []string) {
			ts, _ := strconv.ParseInt(Built, 10, 64)
			fmt.Printf(`Version: %s --HEAD
GitCommit: %s
Built: %s
Kratos Version: %s
Go Version: %s
OS/Arch: %s/%s
`,
				Version,
				GitHash,
				time.Unix(ts, 0).Format("2006-01-02 15:04:05"),
				kratos.Release,
				runtime.Version(),
				runtime.GOOS,
				runtime.GOARCH)
			os.Exit(0)
		},
	}
)
