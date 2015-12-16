package placeholderCount

import "fmt"

//import "log"

func Printf(placeholder string, args ...int) {

}

const (
	myFormatConst       = "%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[9]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d"
	myConcatFormatConst = "%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f " + " %[9]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d"
	myWrongFormatConst  = "%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[9]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d"

	demo1, demo2, demo3 = 1, 1+1-1, "a" + "b" + "c"
)

func _() {
	fmt.Errorf(myConcatFormatConst, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11)

	/*fmt.Fprintf(_, "%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[1]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d",
		1, 2, 3, 4, 5, 6, 7,
	)*/

	/*fmt.Fscanf(_, "%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[1]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d",
		1, 2, 3, 4, 5, 6, 7,
	)
	fmt.Printf("%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[1]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d",
		1, 2, 3, 4, 5, 6, 7,
	)
	fmt.Scanf("%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[1]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d",
		1, 2, 3, 4, 5, 6, 7,
	)
	fmt.Sprintf("%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[1]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d",
		1, 2, 3, 4, 5, 6, 7,
	)
	fmt.Sscanf(_, "%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[1]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d",
		1, 2, 3, 4, 5, 6, 7,
	)
	log.Fatalf("%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[1]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d",
		1, 2, 3, 4, 5, 6, 7,
	)
	log.Panicf("%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[1]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d",
		1, 2, 3, 4, 5, 6, 7,
	)
	log.Printf("%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[9]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d",
		1, 2, 3, 4, 5, 6, 7, 8,
	)

	printf("%d")
	printf("%d", 1)

	printf("%[2]d %[1]d", 1, 2)
	printf("%[2]d %[1]d", 1)
	printf("%[2]d %[3]d", 1, 2)
	printf("%[4]*.f %[3]d", 1, 2)

	fmt.Sprintf("%d")
	fmt.Sprintf("%d", 1, 2)*/
}
