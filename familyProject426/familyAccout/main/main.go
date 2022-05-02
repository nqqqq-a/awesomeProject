package main

import (
	"awesomeProject/familyProject426/familyAccout/utils"
	"fmt"
)

func main() {
	fmt.Println("这是面向对象的方式完成...")
	familyAccount := utils.NewFamilyAccount()
	familyAccount.MainMenu()
}
