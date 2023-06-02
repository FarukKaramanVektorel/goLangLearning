// In this programme, the user is asked to enter the quantity of each fruit and the payment amount is  calculated for each fruit separately. If the user makes an incorrect entry while the user is expected to enter, the programme is terminated.
package main
import "fmt"

func main() {
	fruits := [][]string{
		{"Elma", "9.50"},
		{"Armut", "12.75"},
		{"Muz", "24.20"},
		{"Çilek", "36.80"},
	}
	totalPayment := 0.0
	for _, fruit := range fruits {
		name := fruit[0]
		price := fruit[1]

		var quantity int
		fmt.Printf("Kaç kilo %s almak istiyorsunuz? ", name)
		_, err := fmt.Scanln(&quantity)
		if err != nil {
			fmt.Println("Hatalı bir giriş yaptınız. Program sonlandırılıyor.")
			return
		}

		subtotal := float64(quantity) * parseFloat(price)
		totalPayment += subtotal
	}

	fmt.Printf("Toplam ödemeniz gereken miktar: %.2f TL\n", totalPayment)
}

func parseFloat(s string) float64 {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	if err != nil {
		return 0
	}
	return f
}
