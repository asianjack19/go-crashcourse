package main

import (
	"fmt"
)

// //bikin defer function
// func logging() {
// 	fmt.Println("Selesai memanggil function")
// }
// func runApplication(value int) {
// 	//defer dijalanin even if it has errors
// 	defer logging()
// 	fmt.Println("Run Application")
// 	fmt.Println(10 / value)
// }

//test panic
func endApp() {
	fmt.Println("Aplikasi Selesai")
	//recover buat ngambil text dari panic
	message := recover()
	if message != nil {
		fmt.Println(message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}

	fmt.Println("Aplikasi jalan")
}

func main() {
	// latihan defer
	// runApplication(0)

	runApp(false)
}
