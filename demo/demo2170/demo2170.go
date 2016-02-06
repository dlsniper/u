package demo2170

func _(c interface{}) {
	_ = *(*func(width int) struct{})(nil)
}
