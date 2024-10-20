package main

import (
	algo "algo_start/list"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

const ALGO_HISTORY = "./algo_history.txt"

func main() {
	start_text()
	args := []string{"/c", "start", get_link()}
	exec.Command("cmd", args...).Start()

}
func start_text() {
	talk := "Talk is cheap. Show me the code."
	for _, v := range talk {
		fmt.Printf("%c", v)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println(" ", time.Now().Format("2006-01-02 15:04:05"))

}
func get_link() string {
	var level int
	var link string
	var num int
	level = rand.Intn(10)
	if level <= 2 {
		num = rand.Intn(len(algo.EASY_LIST))
		link = algo.EASY_LIST[num]
		print("【简单】")
	} else if level <= 7 {
		num = rand.Intn(len(algo.MEDIUM_LIST))
		link = algo.MEDIUM_LIST[num]
		print("【中等】")
	} else {
		num = rand.Intn(len(algo.HARD_LIST))
		link = algo.HARD_LIST[num]
		print("【困难】")
	}
	link = strings.Split(strings.Split(link, "(")[1], ")")[0]
	return link
}
