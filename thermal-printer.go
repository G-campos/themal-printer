package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/augustopimenta/escpos"
)

func main() {
	f, err := os.OpenFile("/dev/usb/lp2", os.O_RDWR, 0)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	p := escpos.New(f)
	p.Init()

	p.FontSize(2, 2)
	p.Font(escpos.FontB)
	p.FontAlign(escpos.AlignCenter)
	p.Writeln("Campos Tech")
	p.Feed()

	readFile, err := os.Open("Template.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range lines {
		fmt.Println(line)
		p.FontSize(1, 1)
		p.Font(escpos.FontA)
		p.FontAlign(escpos.AlignLeft)
		p.Writeln(line)

	}

	p.FeedN(3)

	p.FullCut()
}
