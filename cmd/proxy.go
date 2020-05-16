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
	"os"

	"github.com/kris-nova/knobs/stream/rtmp"
	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
)

var ProxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Used to serve a bidirectional server and proxy client",
	Long:  `Used to serve a bidirectional server and proxy client`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Always("Running in [PROXY] mode.")

	},
}

func init() {
	// Commands here
}

func Proxy() {
	// switch{} # cli args
	// srt.New() ?
	// http2.New() ?
	// whatevs.New() ?
	streamer := rtmp.New()
	err := streamer.Stream()
	if err != nil {
		logger.Critical("%+v", err.Error())
		os.Exit(1)
	}
	logger.Always("kthnxbai")
	os.Exit(0)
}
