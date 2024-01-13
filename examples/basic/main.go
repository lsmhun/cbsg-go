package main

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	cbsg "github.com/lsmhun/cbsg-go"
)

func main() {
	cbsgBase := cbsg.NewDefaultCbsgCore()
	dilbert_boss_say(cbsgBase.ShortWorkshop())

}

func dilbert_boss_say(text string) {
	say, err := cowsay.Say(
		text,
		// I added dilbertsboss head to cowsay parts
		cowsay.Type("dilbertsboss"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}
