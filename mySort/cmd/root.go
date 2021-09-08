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
	"bufio"
	"fmt"
	"mySort/pkg/ArrayStrings"
	"os"
	"sort"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mySort",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No such file")
			return
		}

		str, err := readLines(args[0])

		if err != nil {
			fmt.Println("Not found file")
			return
		}

		linesStr := ArrayStrings.NewArrayStrings(str)

		//если флан не выбран, то сортируем стандартной сортировкой
		f := linesStr.StandartSort

		if val, _ := cmd.Flags().GetInt("column"); val > 0 {
			fmt.Println("column", val)
			linesStr.SetSortColumn(val)
		}
		if ok, _ := cmd.Flags().GetBool("unique"); ok {
			fmt.Println("unique")
			linesStr.Unique()
		}
		if ok, _ := cmd.Flags().GetBool("ignore"); ok {
			fmt.Println("is ignore")
			linesStr.IgnoreSpace()
		}

		if ok, _ := cmd.Flags().GetBool("number"); ok {
			f = linesStr.NumberSort
		}

		if ok, _ := cmd.Flags().GetBool("month"); ok {
			fmt.Println("is month")
			f = linesStr.MonthSort
		}

		if ok, _ := cmd.Flags().GetBool("suffix"); ok {
			fmt.Println("is suffix")
		}
		if ok, _ := cmd.Flags().GetBool("check"); ok {
			linesStr.IsSorted(f)
			/*eqVal := make([]string, len(linesStr))
			copy(eqVal, linesStr)
			sort.SliceIsSorted(eqVal, f)
			for i := 0; i < len(linesStr); i++ {

			}*/
			return
		}

		//no flags
		sort.SliceStable(linesStr, f)

		if ok, _ := cmd.Flags().GetBool("reverse"); ok {
			fmt.Println("is rev")

			/*for i, j := 0, linesStr.Len()-1; i < j; i, j = i+1, j-1 {
				linesStr[i], linesStr[j] = linesStr[j], linesStr[i]
			}*/

			/*sort.SliceStable(linesStr, func(i, j int) bool {
				return linesStr[i] >= linesStr[j]
			})*/

		}

		linesStr.PrintIndexValue()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	var count int
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mySort.yaml)")

	rootCmd.Flags().IntVarP(&count, "column", "k", 0, "Sorts by column")
	rootCmd.Flags().BoolP("reverse", "r", false, "Revers sorts")
	rootCmd.Flags().BoolP("number", "n", false, "Sort by number")
	rootCmd.Flags().BoolP("unique", "u", false, "Unique values sort")
	rootCmd.Flags().BoolP("month", "M", false, "Sort month")
	rootCmd.Flags().BoolP("ignore", "b", false, "Ignore tailing space")
	rootCmd.Flags().BoolP("check", "c", false, "Check sort")

	//TODO
	rootCmd.Flags().BoolP("suffix", "H", false, "Check suffix")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//rootCmd.Flags().BoolP("column", "k", false, "Sorts by column")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mySort" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mySort")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sLines = append(sLines, scanner.Text())
	}

	return sLines, scanner.Err()
}
