package main

import ( 
	cram "github.com/aroario2003/cram/cmd"
	gui "github.com/aroario2003/cram/gui"
)

func main() {
	cram.InitCliArgs()
	if cram.GetGui() {
		gui.Show()
	} else {
		cram.Entry()
	}
}
