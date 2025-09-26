package conversion

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Contador() {
	vocales := "aeiou"
	fmt.Print("Introduce cadena de texto: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()
	contador := 0
	frase = strings.ToLower(frase)

	// Recorrer cada carácter
	for _, char := range frase {
		if strings.ContainsRune(vocales, char) {
			contador++
		}
	}
	fmt.Printf("La frase tiene %d vocales\n", contador)
}
