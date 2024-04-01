package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Zipper struct {
	Nodes         []Node
	FreqDic       map[rune]int
	EncriptionMap map[rune][]byte
}

func concatBytes(byteArray []byte) []byte {
	var concatedBytes []byte
	var unit byte = 0

	for i := 0; i < len(byteArray); i++ {
		if i%8 == 0 && i != 0 {
			concatedBytes = append(concatedBytes, unit)
			unit = 0
		}

		unit = 2*unit + byteArray[i]
	}

	concatedBytes = append(concatedBytes, unit)

	return concatedBytes
}

func (z *Zipper) Zip(fileName string) {

	text, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	z.FindEncoding(string(text))

	var byteArray []byte

	for _, c := range string(text) {
		byteArray = append(byteArray, z.EncriptionMap[c]...)
	}

	zippedByteArray := concatBytes(byteArray)

	err = os.WriteFile("zipped-"+fileName, zippedByteArray, 0600)

	if err != nil {
		fmt.Print(err)
	}
}

func (z *Zipper) Unzip(fileName string) {
	text, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	z.DeCode(text)
}

func (z *Zipper) FindEncoding(text string) {
	z.FindFreq(text)
	z.CreateNodes()
	z.CreateTree()
	z.EncriptionMap = make(map[rune][]byte)
	var code []byte
	z.Nodes[0].Encode(code, z.EncriptionMap)
}

func (z *Zipper) FindFreq(text string) {
	z.FreqDic = make(map[rune]int)
	for _, c := range text {
		z.FreqDic[c]++
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

func (z *Zipper) DeCode(text []byte) {
	var unit []byte
	var sb strings.Builder
	bits := convertBack(text)
	for _, c := range bits {
		unit = append(unit, c)
		if key, ok := mapkey(z.EncriptionMap, unit); ok {
			sb.WriteRune(key)
			unit = nil
		} else {
			continue
		}
	}
	err := os.WriteFile("original.txt", []byte(sb.String()), 0644)
	if err != nil {
		fmt.Print(err)
	}
}

func convertBack(text []byte) []byte {
	var totalByteArray []byte
	var byteArray []byte
	var reminder int
	for n, i := range text {
		reminder = int(i)
		for len(byteArray) < 8 && n != len(text)-1 || (reminder > 0) {
			bit := reminder % 2
			reminder = reminder / 2
			byteArray = append([]byte{byte(bit)}, byteArray...)
		}
		totalByteArray = append(totalByteArray, byteArray...)
		byteArray = nil
	}

	return totalByteArray
}

func mapkey(m map[rune][]byte, value []byte) (key rune, ok bool) {
	for k, v := range m {
		if bytes.Equal(v, value) {
			key = k
			ok = true
			return
		}
	}
	return
}
