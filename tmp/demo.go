package main

type (
	Item struct {
		Name string
	}

	Menu []Item
)

const demo = 3

func main() {
	_ = Menu{
		{Name: "h1ome"}, // flagged as unnamed field initialization error but it's correct
	}
}
