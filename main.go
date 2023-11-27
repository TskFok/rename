package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("输入季度")
	season, _ := reader.ReadString('\n')
	season = strings.TrimSuffix(season, "\n")

	sInt, _ := strconv.Atoi(season)
	if sInt < 10 {
		season = "0" + season
	}

	fmt.Println("输入路径")
	dir, _ := reader.ReadString('\n')
	dir = strings.TrimSuffix(dir, "\n")

	fmt.Println("输入关键词")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSuffix(keyword, "\n")

	fmt.Println("集数位置")
	position, _ := reader.ReadString('\n')
	position = strings.TrimSuffix(position, "\n")
	positionInt, _ := strconv.Atoi(position)

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
			}
		}

		lInt, _ := strconv.Atoi(l)
		if lInt < 10 {
			l = "0" + l
		}

		newPath := dir + "/" + "S" + season + "E" + l + ext
		err := os.Rename(dir+"/"+file.Name(), newPath)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(newPath)
	}
}
