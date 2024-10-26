package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := DeleteSliceElementV2(1, s1)
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	// FIXME: 這邊想請教老師，打印出來的結果會是
	// s3: [foo bar baz qux], len: 4, cap: 4
	// s3: [foo bar qux qux], len: 4, cap: 4 => FIXME: 原始變量的值被修改了，這個該怎麼解？
	// s4: [foo bar qux], len: 3, cap: 4
	s3 := []string{"foo", "bar", "baz", "qux"}
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
	s4 := DeleteSliceElementV3(2, s3)
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
	fmt.Printf("s4: %v, len: %d, cap: %d \n", s4, len(s4), cap(s4))

	s5 := []string{"foo", "bar", "baz", "qux"}
	fmt.Printf("s5: %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))
	s6 := DeleteSliceElementV4[string](2, s5)
	fmt.Printf("s5: %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))
	fmt.Printf("s6: %v, len: %d, cap: %d \n", s6, len(s6), cap(s6))
}
