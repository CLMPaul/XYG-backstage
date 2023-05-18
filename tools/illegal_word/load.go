package illegal_word

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"xueyigou_demo/global"
	_ "xueyigou_demo/global"
)

func Load() {
	content, err := os.Open("自定义敏感词.txt")
	// handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer content.Close()

	fileScanner := bufio.NewScanner(content)
	list := make([]string, 0)
	// read line by line
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		list = append(list, fileScanner.Text())
	}
	fmt.Println(list)
	//
	global.IllegalWords.SetKeywords(list)
	global.IllegalWords.Save("sensitive.dat")
}
