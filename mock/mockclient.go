package main

import (
	p "mir-go/proto"
	cp "mir-go/proto/client"
)

func main() {
	// client version
	pkg := &p.Packet{false, cp.CLIENT_VERSION, &cp.ClientVersion{}}
	bytes := pkg.ToBytes()

	// login

	// TODO new character

	// start game

	// refine cancel ??

	//// in game

	// walk

	// direction

}
