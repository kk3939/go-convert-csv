package cmd

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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
		data := readJson()
		convertCsv(data)
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
func readJson() []JsonType {
	// main.go を起点としたpath
	data, err := ioutil.ReadFile("./json_dir/sample.json")

	// 読み込みエラーがあったら、exit
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var dt []JsonType

	// data をバイトデータから変換
	// 第二引数にポインタを渡すのが注意
	if err := json.Unmarshal(data, &dt); err != nil {
		fmt.Printf("Could not read data. %v", err)
	}

	return dt
}

// コンバートする
func convertCsv(data []JsonType) {
	csvf, err := os.Create("./csv_dir/converted.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvf.Close()

	writer := csv.NewWriter(csvf)
	for i, d := range data {
		var row []string

		if i == 0 {
			row = append(row, "Nation")
			row = append(row, "Region")
			row = append(row, "Capital")
			writer.Write(row)
		}
		row = append(row, d.Nation)
		row = append(row, d.Region)
		row = append(row, d.Capital)
		fmt.Println(row)
		writer.Write(row)
	}
	writer.Flush()
}

//Jsonの型定義
type JsonType struct {
	Nation  string
	Region  string
	Capital string
}
