/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "批量重命名电视剧名称",
	Long:  `批量重命名电视剧名称到可刮削的名称`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("输入季度")
		season, _ := reader.ReadString('\n')
		season = strings.TrimSuffix(season, "\n")

		sInt, _ := strconv.Atoi(season)
		if sInt < 10 {
			season = "0" + season
		}

		fmt.Println("输入文件中包含的关键词")
		keyword, _ := reader.ReadString('\n')
		keyword = strings.TrimSuffix(keyword, "\n")

		fmt.Println("集数所在位置,从0开始")
		position, _ := reader.ReadString('\n')
		position = strings.TrimSuffix(position, "\n")
		positionInt, _ := strconv.Atoi(position)

		fmt.Println("倒退集数,不倒退填写0")
		back, _ := reader.ReadString('\n')
		back = strings.TrimSuffix(back, "\n")
		backInt, _ := strconv.Atoi(back)

		fmt.Println(season, dir, keyword, position, back)

		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("读取目录失败:", err)
			return
		}

		for _, file := range files {
			if !strings.Contains(file.Name(), keyword) {
				continue
			}
			fmt.Println(file.Name())
			ext := path.Ext(dir + "/" + file.Name())

			regex := regexp.MustCompile(`\d+`)
			matches := regex.FindAllString(file.Name(), -1)

			l := "0"
			for i, match := range matches {
				if i == positionInt {
					l = match

					lInt, _ := strconv.Atoi(l)
					lInt = lInt - backInt
					l = strconv.Itoa(lInt)

					if lInt < 10 {
						l = "0" + l
					}
				}
			}

			newPath := dir + "/" + "S" + season + "E" + l + ext
			err := os.Rename(dir+"/"+file.Name(), newPath)

			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println(newPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
