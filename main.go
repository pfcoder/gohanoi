package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	PA = iota
	PB
	PC
)

var steps = 0

var layerCnt = 3 // Default 3 layers
var pH [3][]int

func main() {
	args := os.Args
	for index, v := range args {
		if v == "-n" {
			// Next arg is layer size
			layerCnt, _ = strconv.Atoi(args[index+1])
		}
	}
	pH[PA] = make([]int, layerCnt)
	pH[PB] = make([]int, layerCnt)
	pH[PC] = make([]int, layerCnt)
	for i := 0; i < layerCnt; i++ {
		pH[PA][i] = i*2 + 2
	}

	outputStatus()
	move(PA, PC, PB, layerCnt)
	fmt.Printf("Move done, steps:%v\n", steps)
}

func avaTpos(aIndex int) int {
	for i := len(pH[aIndex]) - 1; i >= 0; i-- {
		if pH[aIndex][i] == 0 {
			return i
		}
	}
	return -1
}

func avaSpos(aIndex int) int {
	for i := 0; i < len(pH[aIndex]); i++ {
		if pH[aIndex][i] != 0 {
			return i
		}
	}
	return -1
}

func moveOne(source int, target int) {
	posT := avaTpos(target)
	posS := avaSpos(source)
	pH[target][posT] = pH[source][posS]
	pH[source][posS] = 0
	steps++
	outputStatus()
}

func move(s int, t int, tmp int, moveSize int) {
	//fmt.Printf("Enter move, s:%v t:%v tmp:%v size:%v\n", s, t, tmp, moveSize)
	if moveSize == 1 {
		moveOne(s, t)
		return
	}

	// Step 1 move top to tmp
	move(s, tmp, t, moveSize-1)

	// Step 2 move bottom to target
	moveOne(s, t)

	// Step 3 move top from tmp to target
	move(tmp, t, s, moveSize-1)
}

func printLayer(baseW int, w int) {
	if w > 0 {
		tOffset := (baseW - w) / 2
		fmt.Printf(strings.Repeat(" ", tOffset))
		fmt.Printf(strings.Repeat("▓", w))
		fmt.Printf(strings.Repeat(" ", tOffset))
	} else {
		fmt.Printf(strings.Repeat(" ", baseW))
	}
}

func clearScreen() {
	var clrCmd string
	if runtime.GOOS == "windows" {
		clrCmd = "cls"
	} else {
		clrCmd = "clear"
	}

	cmd := exec.Command(clrCmd)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printOffset() {
	fmt.Printf("    ")
}

func outputStatus() {
	// Visual output
	clearScreen()

	for i := 0; i < 15; i++ {
		fmt.Println("")
	}
	for i := 0; i < layerCnt; i++ {
		printOffset()

		// Tower A
		printLayer(layerCnt*2, pH[PA][i])

		// Tower B
		printLayer(layerCnt*2, pH[PB][i])

		// Tower C
		printLayer(layerCnt*2, pH[PC][i])

		fmt.Println("")
	}

	// Print base line
	printOffset()

	fmt.Printf(strings.Repeat("─", layerCnt*2*3) + "\n")
	time.Sleep(1 * time.Second)
}
