package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "go-convert-csv",
	Short: "convert json to csv",
	Long:  `If you put json_file in json_dir, this program convert the file to csv in csv_dir.`,
	Args: func(cmd *cobra.Command, args []string) error {
		// 引数があるとエラーにする
		if len(args) > 0 {
			e_msg := "This program doesn't need args!"
			return errors.New(e_msg)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		readJson()
		convertCsv()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-convert-csv.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-convert-csv")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// json を読み込む
func readJson() {
	// main.go を起点としたpath
	data, err := ioutil.ReadFile("./json_dir/sample.json")

	// 読み込みエラーがあったら、exit
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	dt := JsonType{}

	// data をバイトデータから変換
	if err := json.Unmarshal(data, &dt); err != nil {
		fmt.Printf("Could not read data. %v", err)
	}

	fmt.Println(dt)

}

func convertCsv() {
	fmt.Println("convert CSV is not set up.")
}

//Jsonの型定義
type JsonType struct {
	Nation  string `json:"nation"`
	Region  string `json:"region"`
	Capital string `json:"capital"`
}
