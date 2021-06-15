package main

// Serial confinement is when a variable is shared in a pipeline
// through a channel and each stage of the pipeline
// refrains from acessing the variable after sending it to
// the next stage.

// NOTE:
// Serial confinement is not enforced by the compiler,
// there's nothing stopping users from mutating variables anywhere.
// This problem wouldn't exist if Go data structures were immutable.

type Cake struct {
	state string
}

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer never touchs this cake again
	}
}

func main() {

}
