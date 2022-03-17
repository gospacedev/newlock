package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PasswordGenerator(passwordLength int) string{
	lowCase := "azwsxedcrfvtgbyhnujmikolp"

	highCase := "AZWSXEDCRFVTGBYHNUJMIKOLP"

	nums := "0123456789"

	specChar := "!@#$%^&*()_+[]{}"

	password := ""

	for i := 0; i < passwordLength; i++ {
		rand.Seed(time.Now().UnixNano())
		
        randNum := rand.Intn(4)

		fmt.Println(randNum)

		switch randNum {
		case 0:
			rand.Seed(time.Now().UnixNano())

			randNum := rand.Intn(len(lowCase))

			password = password + string(lowCase[randNum])

		case 1:
			rand.Seed(time.Now().UnixNano())

			randNum := rand.Intn(len(highCase))

			password = password + string(highCase[randNum])

		case 2:
			rand.Seed(time.Now().UnixNano())

			randNum := rand.Intn(len(nums))

			password = password + string(nums[randNum])

		case 3:
			rand.Seed(time.Now().UnixNano())

			randNum := rand.Intn(len(specChar))

			password = password + string(specChar[randNum])

		}

	}

	fmt.Println(password)

	return password
}