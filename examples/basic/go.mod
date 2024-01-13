module github.com/lsmhun/cbsg-go/example

go 1.19

require github.com/Code-Hex/Neo-cowsay/v2 v2.0.4

replace github.com/Code-Hex/Neo-cowsay/v2 => github.com/lsmhun/Neo-cowsay/v2 v2.0.5


require (
	github.com/Code-Hex/go-wordwrap v1.0.0 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
)

require github.com/lsmhun/cbsg-go v0.2.0

replace github.com/lsmhun/cbsg-go => ../../
