package issue


func _() {
	a := Demo{
		Field: 1, // Refactor here
	}
	a.Field2 = 1 // Missing quick action here
	_ = a
}