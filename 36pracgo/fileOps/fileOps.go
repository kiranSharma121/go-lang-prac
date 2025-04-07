package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func WriteDataToFile(balance float64, filename string) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(filename, []byte(balanceTxt), 0664)
}
func GetDataFromFile(filename string) float64 {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000
	}
	datatxt := string(data)
	amount, err := strconv.ParseFloat(datatxt, 64)
	if err != nil {
		fmt.Println("Error in parse the data")
	}
	return amount

}
