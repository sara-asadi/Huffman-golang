package main

import (
	"fmt"
	"os"
)

type Zipper struct {
	Nodes         []Node
	FreqDic       map[string]int
	EncriptionMap map[string]string
}

func (z *Zipper) FindFreq(text string) {
	z.FreqDic = make(map[string]int)
	for _, c := range text {
		_, ok := z.FreqDic[string(c)]
		if !ok {
			z.FreqDic[string(c)] = 1
		} else {
			z.FreqDic[string(c)] += 1
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
	z.EncriptionMap = make(map[string]string)
	z.Nodes[0].Encode("", z.EncriptionMap)
}

func (z *Zipper) Zip(fileName string) {

	text, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	z.FindEncoding(string(text))

	var encryptedText string

	for _, c := range text {
		encryptedText += z.EncriptionMap[string(c)]
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
	var decryptedText string
	for _, c := range text {
		unit += string(c)
		if key, ok := mapkey(z.EncriptionMap, unit); ok {
			decryptedText += key
			unit = ""
		} else {
			continue
		}
	}
	err := os.WriteFile("original.txt", []byte(decryptedText), 0644)
	if err != nil {
		fmt.Print(err)
	}
}

func mapkey(m map[string]string, value string) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
