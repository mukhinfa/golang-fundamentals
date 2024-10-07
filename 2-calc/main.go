package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var menu = map[string]func([]int){
	"AVG": calculateAverage,
	"SUM": calculateSum,
	"MED": calculateMedian,
}

func main() {
	values := getValues()
	operation := getOperation()
	menuFunc := menu[operation]
	menuFunc(values)
}

func getOperation() string {
	var v string
	fmt.Print("Введите операцию: ")
	for {
		fmt.Scan(&v)
		if v == "AVG" || v == "SUM" || v == "MED" {
			return v
		}
		fmt.Printf("У нас нет операции %s\n", v)
		fmt.Print("Введите операцию еще раз: ")
	}
}

func getValues() []int {
	fmt.Print("Введите значения через запятую ")

	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	stringValues := strings.Split(strings.ReplaceAll(line, " ", ""), ",")
	values := make([]int, 0, len(stringValues))
	for _, v := range stringValues {
		num, err := strconv.Atoi(v)
		if err == nil {
			values = append(values, num)
		}
	}
	return values
}

func calculateAverage(arr []int) {
	var sum int
	for _, v := range arr {
		sum += v
	}
	fmt.Printf("Среднее значение: %v", float64(sum/len(arr)))
}

func calculateSum(arr []int) {
	var sum int
	for _, v := range arr {
		sum += v
	}
	fmt.Printf("Сумма значений: %v", sum)
}

func calculateMedian(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	if len(arr)%2 != 0 {
		fmt.Printf("Медианное значение: %v", arr[len(arr)/2])
	}
	fmt.Printf("Медианное значение: %v", float64(arr[len(arr)/2]+arr[len(arr)/2-1])/2)
}
