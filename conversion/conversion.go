package conversion

import (
	"fmt"
	"strings"
)

func Conversion() {
	var dolares float64
	var moneda string

	// Tasas de conversión
	tasaEUR := 0.95     // Euro
	tasaGBP := 0.82     // Libra esterlina
	tasaKRW := 1375.00  // Won surcoreano
	tasaBTC := 0.000015 // Bitcoin
<<<<<<< HEAD
	fmt.Println("Ingrese la cantadidad en dólares: ")
=======
	fmt.Print("Ingrese la cantidad en dólares: ")
>>>>>>> 45ed494 (Agregar paquetes conversion y contadorVocales y actualizar main.go)
	fmt.Scanln(&dolares)
	fmt.Print("Conversiones disponibles EUR/GBP/KRW/BTC: ")
	fmt.Scanln(&moneda)
<<<<<<< HEAD
	moneda = strings.ToUpper(moneda)
	dolarEuro := dolares * tasaEUR
	dolarLibras := dolares * tasaGBP
	dolarWon := dolares * tasaKRW
	dolarBTC := dolares * tasaBTC
	switch moneda {
	case "EUR":
		fmt.Println("La conversión es de", dolarEuro)
	case "GBP":
		fmt.Println("La conversión es de", dolarLibras)
=======
	moneda = strings.ToUpper(moneda) // para evitar problemas de mayúsculas/minúsculas

	switch moneda {
	case "EUR":
		fmt.Println("La conversión es de", dolares*tasaEUR, "Euros")
	case "GBP":
		fmt.Println("La conversión es de", dolares*tasaGBP, "Libras Esterlinas")
>>>>>>> 45ed494 (Agregar paquetes conversion y contadorVocales y actualizar main.go)
	case "KRW":
		fmt.Println("La conversión es de", dolares*tasaKRW, "Won Surcoreano")
	case "BTC":
		fmt.Println("La conversión es de", dolares*tasaBTC, "Bitcoin")
	default:
		fmt.Println("Moneda no reconocida")
	}
}
