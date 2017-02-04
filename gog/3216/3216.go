package main

const str =
	`TempDir creates a new temporary directory in the directory dir with a name beginning with prefix and returns
	the path of the new directory. If dir is the empty string, TempDir uses the default directory for temporary files (see
	os.TempDir). Multiple programs calling TempDir simultaneously will not choose the same directory. It is the caller's
	responsibility to remove the directory when no longer needed.
	`

func main() {
	_ = str
}
