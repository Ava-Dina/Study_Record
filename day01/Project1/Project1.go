package main

import (
	"fmt"
	"math/rand"
	"time"
)

//项目1:猜谜游戏
func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber) //答案

	fmt.Println("Please input your guess")
	//reader := bufio.NewReader(os.Stdin)						** origin version
	for {
		var guess int                         //change
		_, err := fmt.Scanf("%d\r\n", &guess) //change
		//input, err := reader.ReadString('\n')					** origin version
		if err != nil {
			fmt.Println("An error occured while reading input.Please try again", err)
			continue
		}
		//** origin version
		//input = strings.TrimSuffix(input, "\r\n") //这里尤其注意在windows平台换行符为\r\n,Linux和MacOS平台则为\n
		//
		//guess, err := strconv.Atoi(input)
		//if err != nil {
		//	fmt.Println("Invalid input. Please enter an integer value")
		//	continue
		//}												//** origin version
		fmt.Println("Your guess is ", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number.Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number.Please try again")
		} else {
			fmt.Println("Correct,you Legend!")
			break
		}
	}
}
