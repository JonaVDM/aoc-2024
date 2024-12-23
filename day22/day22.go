package day22

import "strconv"

func Run(data []string) [2]interface{} {
	var a int
	history := make([][]int, len(data))

	for i, d := range data {
		n, _ := strconv.Atoi(d)
		score, hist := GetScore(n, 2000)
		n += score
		history[i] = hist
	}

	return [2]interface{}{
		a,
		GetBestDeal(history),
	}
}

func GetScore(start, amount int) (int, []int) {
	secret := start
	history := make([]int, amount)
	for i := range amount {
		var tmp int

		tmp = secret * 64
		secret = Mix(tmp, secret)
		secret = Prune(secret)

		tmp = secret / 32
		secret = Mix(tmp, secret)
		secret = Prune(secret)

		tmp = secret * 2048
		secret = Mix(tmp, secret)
		secret = Prune(secret)

		history[i] = GetLastDigit(secret)
	}
	return secret, history
}

func Mix(tmp, secret int) int {
	return tmp ^ secret
}

func Prune(secret int) int {
	return secret % 16777216
}

func GetLastDigit(num int) int {
	return num % 10
}

func GetBestDeal(prices [][]int) int {
	for i := 5; i < len(prices[0]); i++ {
	}
	return 0
}
