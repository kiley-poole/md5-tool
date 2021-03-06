package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Marvel API Hash Gen")
	fmt.Println("---------------------")

	fmt.Println("Enter your public key: ")
	pk, _ := reader.ReadString('\n')
	pk = strings.Replace(pk, "\n", "", -1)

	fmt.Println("Enter your private key: ")
	pvk, _ := reader.ReadString('\n')
	pvk = strings.Replace(pvk, "\n", "", -1)

	r := strconv.Itoa(rand.Int())

	h := md5.New()

	h.Write([]byte(r + pvk + pk))

	bs := h.Sum(nil)

	fmt.Println("Your url is:")
	fmt.Printf("?ts=%s&apikey=%s&hash=%x\n", r, pk, bs)

}
