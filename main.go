package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
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
	dir := getDir()

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
}

func getDir() string {
	// 开始捕获键盘输入
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	dir := ""
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyCtrlC {
			break
		}

		if key == keyboard.KeyEnter {
			return dir
		}

		if key == keyboard.KeyTab {
			fmt.Println("")
			files, err := os.ReadDir(dir)

			if err != nil {
				fmt.Println("读取目录失败:", err)
				dir = ""
				continue
			}

			list := ""

			for _, file := range files {
				list += file.Name() + " "
			}

			if len(files) == 1 {
				dir += "/" + files[0].Name()
				fmt.Println(dir)
			} else {
				dir = ""
				fmt.Println(list)
			}
			continue
		}

		fmt.Printf("%c", char)
		dir += fmt.Sprintf("%c", char)
	}

	return ""
}
