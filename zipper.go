package main

import (
	"fmt"
	"os"
	"strings"
)

type Zipper struct {
	Nodes         []Node
	FreqDic       map[rune]int
	EncriptionMap map[rune]string
}

func (z *Zipper) FindFreq(text string) {
	z.FreqDic = make(map[rune]int)
	for _, c := range text {
		_, ok := z.FreqDic[c]
		if !ok {
			z.FreqDic[c] = 1
		} else {
			z.FreqDic[c] += 1
		}
	}
}

func (z *Zipper) CreateNodes() {
	for key, freq := range z.FreqDic {
		z.Nodes = append(z.Nodes, *NewNode(key, freq))
	}
}

func (z *Zipper) FindMinFreq() Node {
	minNode := z.Nodes[0]
	index := 0
	for i, n := range z.Nodes {
		if minNode.freq > n.freq {
			minNode = n
			index = i
		}
	}
	z.Nodes = append(z.Nodes[:index], z.Nodes[index+1:]...)

	return minNode
}

func (z *Zipper) CreateTree() {

	for len(z.Nodes) > 1 {
		minNode1 := z.FindMinFreq()
		minNode2 := z.FindMinFreq()

		newNode := Add(&minNode1, &minNode2)

		z.Nodes = append(z.Nodes, *newNode)
	}
}

func (z *Zipper) FindEncoding(text string) {
	z.FindFreq(text)
	z.CreateNodes()
	z.CreateTree()
	z.EncriptionMap = make(map[rune]string)
	z.Nodes[0].Encode("", z.EncriptionMap)
}

func (z *Zipper) Zip(fileName string) {

	text, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	z.FindEncoding(string(text))

	var encryptedText []byte

	var block byte
	var size int = 0
	for _, c := range text {
		encryptedText += z.EncriptionMap[rune(c)]
	}

	err = os.WriteFile("zipped-"+fileName, []byte(encryptedText), 0644)
	if err != nil {
		fmt.Print(err)
	}
}

func (z *Zipper) Unzip(fileName string) {
	text, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	z.DeCode(string(text))

}

func (z *Zipper) DeCode(text string) {
	var unit string

	var sb strings.Builder

	for _, c := range text {
		unit += string(c)
		if key, ok := mapkey(z.EncriptionMap, unit); ok {
			sb.WriteRune(key)
			unit = ""
		} else {
			continue
		}
	}
	err := os.WriteFile("original.txt", []byte(sb.String()), 0644)
	if err != nil {
		fmt.Print(err)
	}
}

func mapkey(m map[rune]string, value string) (key rune, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
