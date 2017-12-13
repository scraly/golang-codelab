package basics

func MustReturnAnInteger() int {
	return 5
}

func MustReturnTheSum(a, b int) int {
	return a + b
}

func MustReturnAFloat() float64 {
	return 0.2
}

func MustBeTrue() bool {
	return true
}

func MustBeFalse() bool {
	return false
}

func MustBeTrueIfThereAreTenEggsOrMore(numberOfEggs int) bool {
	if numberOfEggs >= 10 {
		return true
	} else {
		return false
	}

}

func MustReturnHello() string {
	return "Hello"
}

// Here you should only add new lines without changing the existing ones
func MustReturnWorld() string {
	var who = "world"
	return who
}

// Here you should only add new lines without changing the existing ones
func MustSetToTrueAndReturn() bool {
	var changeMe bool
	changeMe = true
	return changeMe
}
