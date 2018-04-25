// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ideaCmd represents the idea command
var ideaCmd = &cobra.Command{
	Use:   "idea",
	Short: "Install idea",
	Long: `Download Intellij Idea Ultimate from jetbrains.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("idea called")
	},
}

func init() {
	installCmd.AddCommand(ideaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ideaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	ideaCmd.Flags().StringP("install-path", "p", "/opt", "What dir to install Intellij Idea in.")
	ideaCmd.Flags().BoolP("create-symlink", "s", true, "Create symlink in install folder.")
	ideaCmd.Flags().StringP("symlink-name", "n", "idea", "Name of the symlink created with -s.")
	ideaCmd.Flags().StringP("version", "v", "2018.1.2", "What version of Intellij Idea to install.")
}
