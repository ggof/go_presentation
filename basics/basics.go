package basics

import (
	"fmt"
	"math/rand"
)

// definition de variable globale de package
var integer int

// fonction privée
func privateFunc() {}

// fonction publique
func PublicFunc() {}

// les fonctions peuvent retourner plus d'un arument
func SayHello(name string) (string, error) {
	// condition
	if name == "Felix Thibault" {
		return "", fmt.Errorf("ne merite pas de bonjour")
	}

	return fmt.Sprintf("BONJOUUUR %s", name), nil
}

func Boucles() {
	array := make([]int, 256)

	// boucle par index
	for i := range array {
		array[i] = rand.Int()
	}

	// boucle par valeur
	for index, value := range array {
		fmt.Printf("value of index %d is %d\n", index, value)
	}
}

// réception par pointeur
func ByReference(ref *string) {
	// les pointeurs peuvent être nuls!
	if ref != nil {
		// do something!
	}
}

func SendByReference() {
	myString := "Bonjour!"
	// & pour envoyer par reference (ATTENTION!)
	ByReference(&myString)
}
