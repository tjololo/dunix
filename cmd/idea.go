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
	"github.com/spf13/cobra"
	"fmt"
	"github.com/tjololo/go-learning/pkg/fileutils"
	"path/filepath"
	"os"
	"github.com/tjololo/go-learning/pkg/idea"
)

var installPath string
var createSymlink bool
var symlinkName string
var version string
// ideaCmd represents the idea command
var ideaCmd = &cobra.Command{
	Use:   "idea",
	Short: "Install idea",
	Long:  `Download Intellij Idea Ultimate from jetbrains.com`,
	Run:   installIdea,
}

func init() {
	installCmd.AddCommand(ideaCmd)
	ideaCmd.Flags().StringVarP(&installPath, "install-path", "p", "/opt", "What dir to install Intellij Idea in.")
	ideaCmd.Flags().BoolVarP(&createSymlink, "create-symlink", "s", true, "Create symlink in install folder.")
	ideaCmd.Flags().StringVar(&symlinkName, "symlink-name", "idea", "Name of the symlink created with -s.")
	ideaCmd.Flags().StringVarP(&version, "version", "v", "2018.1.2", "What version of Intellij Idea to install.")
}

func installIdea(cmd *cobra.Command, args []string) {
	downloadUrl := idea.GetDownloadURI(version)
	downloadIdeaTarTo := "/tmp/idea.tar.gz"
	defer os.Remove(downloadIdeaTarTo)
	if err := idea.DownloadFile(downloadIdeaTarTo, downloadUrl); err != nil {
		fmt.Printf("failed to download idea from URL: %s to folder %s\n%v", downloadUrl, installPath, err)
		return
	}
	fmt.Printf("Extracting to: %s\n", installPath)

	if err := fileutils.Untar(downloadIdeaTarTo, installPath); err != nil {
		fmt.Printf("failed to untar: %s\n%v", downloadIdeaTarTo, err)
		return
	}
	sourceSymlink := filepath.Join(installPath, fileutils.GetFolder(downloadIdeaTarTo))
	symlink := filepath.Join(installPath, symlinkName)
	fmt.Printf("Creating symlink %s -> %s\n", sourceSymlink, symlink)
	if err := fileutils.CreateUpdateSymlink(sourceSymlink, symlink); err != nil {
		fmt.Printf("failed to create symlink: %s\n%v", downloadIdeaTarTo, err)
	}
}
