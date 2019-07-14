/**
 * @Author: harshkhajuria
 * @Date:   01-Jul-2019 06:58:28 am
 * @Email:  khajuriaharsh729@gmail.com
 * @Filename: potato.go
 * @Last modified by:   harshkhajuria
 * @Last modified time: 15-Jul-2019 12:40:05 am
 */

package main

import (
	"fmt"
	"github.com/horcrux2301/Potato/src/potato"
)

func main() {
	if err := potato.Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("running here")
}
