package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(vcode)
}
