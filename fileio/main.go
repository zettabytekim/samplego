package main

import (
	"io"
	"io/ioutil"
	"os"
)

func filecopy() {
	bytes, err := ioutil.ReadFile("../files/1.txt")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("../files/1-copy.txt", bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	fi, err := os.Open("../files/1.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fo, err := os.Create("../files/2.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if cnt == 0 {
			break
		}

		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}

	filecopy()
}
