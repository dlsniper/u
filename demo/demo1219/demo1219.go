package main

func main() {
	var c, d = func() (int, int, int) {return 1, 1, 1}()
	_, _ = c, d
}
