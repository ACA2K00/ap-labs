# Build & Run
In order to Build & Run the project, once downloaded, write the following commands in a terminal

$ go get github.com/hajimehoshi/ebiten/v2
$ go build main.go
$ ./main 30 7

This will install the ebiten library required for the program, build the binary and execute it with the parameters of 30 food items and 7 enemies on the field,
these latter inputs can be changed to once's liking.

An alternative to this is the following command

$ make run

Which will make use of the Makefile commands to execute the prior commands at once, but the food and item amount can't be modified.
