package conversion

import "fmt"

func conversion() {
	var dolares float64
	var moneda string
	tasaEUR := 0.95     // Euro
	tasaGBP := 0.82     // Libra esterlina
	tasaKRW := 1375.00  // Won surcoreano
	tasaBTC := 0.000015 // Bitcoin
	dolarEuro := dolares * tasaEUR
	dolarLibras := dolares * tasaGBP
	dolarWon := dolares * tasaKRW
	dolarBTC := dolares * tasaBTC
	fmt.Println("Ingrese la cantadidad en dólares: ")
	fmt.Scanln(&dolares)
	fmt.Println("Conversiones disponibles EUR/GBP/KRW/BTC : ")
	fmt.Scanln(&moneda)

	switch moneda {
	case "EUR":
		fmt.Println("La conversión es de", dolarEuro)
	case "GBR":
		fmt.Println("La conversión es de", dolarLibras)
	case "KRW":
		fmt.Println("La conversión es de", dolarWon)
	case "BTC":
		fmt.Println("La conversión es de", dolarBTC)
	}

}
