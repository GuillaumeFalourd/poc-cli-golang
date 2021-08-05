/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Command to Login (Short)",
	Long:  `Command to Login (Long)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Login called")
		// get the flag value, its default value is false

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if username != "" && password != "" {
			login(username, password)
		} else {
			fmt.Println("Flags not informed. Please check the command usage with `my-cli login --help`")
		}

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "Add username value")
	loginCmd.Flags().StringP("password", "p", "", "Add password value")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func login(username string, password string) {
	fmt.Println("Login Method called")
	if username == password {
		fmt.Println("LOGIN SUCCESS")
	} else {
		fmt.Println("INVALID PASSWORD")
	}
}
