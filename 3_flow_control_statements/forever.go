package main

func main() {
	count := 1
	for {
		print(count)
		count += 1
		// It's better put a break instead, haha!
		break
	}
}
