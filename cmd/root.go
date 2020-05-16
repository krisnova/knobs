/*
Copyright © 2020 Kris Nóva <kris@nivenly.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "fmt"
  "github.com/kris-nova/logger"
  "github.com/spf13/cobra"
  "os"
)


var RootCmd = &cobra.Command{
  Use:   "knobs",
  Short: "A handy dandy media streaming server, client, and proxy",
  Long: `Use this for stream things in Kubernetes`,
  Run: func(cmd *cobra.Command, args []string) {



    logger.Always("hi")
    os.Exit(1)
  },
}

func Execute() {
  if err := RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
// Commands here
}


