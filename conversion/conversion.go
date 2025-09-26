package conversion

import "fmt"

func Conversion() {
	var dolares float64
	var moneda string
	tasaEUR := 0.95     // Euro
	tasaGBP := 0.82     // Libra esterlina
	tasaKRW := 1375.00  // Won surcoreano
	tasaBTC := 0.000015 // Bitcoin

	fmt.Println("Ingrese la cantidad en dólares: ")
	fmt.Scanln(&dolares)
	fmt.Println("Conversiones disponibles EUR/GBP/KRW/BTC : ")
	fmt.Scanln(&moneda)

	switch moneda {
	case "EUR":
		fmt.Println("La conversión es de", dolares*tasaEUR)
	case "GBP":
		fmt.Println("La conversión es de", dolares*tasaGBP)
	case "KRW":
		fmt.Println("La conversión es de", dolares*tasaKRW)
	case "BTC":
		fmt.Println("La conversión es de", dolares*tasaBTC)
	default:
		fmt.Println("Moneda no válida")
	}
}
