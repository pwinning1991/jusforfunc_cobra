/*
Copyright © 2019 Philip Winnington

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
	"os"

	"github.com/pwinning1991/todo"
	"github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A todo application",
	Long:  `A todo application that can list and add items to a todo list`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var client todo.TasksClient

func init() {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v", err)
		os.Exit(1)
	}
	client = todo.NewTasksClient(conn)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
