package main

/*
*
实现删除切片特定下标元素的方法。

要求一：能够实现删除操作就可以。
要求二：考虑使用比较高性能的实现。
要求三：改造为泛型方法
要求四：支持缩容，并旦设计缩容机制。
*/

func DeleteSliceElementV1(index int, elements []int) []int {
	elements = append(elements)

	for i := 0; i < len(elements)-1; i++ {
		if i >= index {
			elements[i] = elements[i+1]
		}
	}

	return elements[:len(elements)-1]
}

func DeleteSliceElementV2(index int, elements []int) []int {
	newElements := make([]int, len(elements)-1)
	for i := 0; i < len(elements); i++ {
		if i < index {
			newElements[i] = elements[i]
		} else if i > index {
			newElements[i-1] = elements[i]
		}
	}
	return newElements
}

func DeleteSliceElementV3[T any](index int, elements []T) []T {
	elements = append(elements)

	for i := 0; i < len(elements)-1; i++ {
		if i >= index {
			elements[i] = elements[i+1]
		}
	}

	return elements[:len(elements)-1]
}

// FIXME: 這邊雖然能縮容，但重新賦值感覺性能不會太好，請問老師有什麼建議嗎？
func DeleteSliceElementV4[T any](index int, elements []T) []T {
	newElements := make([]T, len(elements)-1)
	for i := 0; i < len(elements); i++ {
		if i < index {
			newElements[i] = elements[i]
		} else if i > index {
			newElements[i-1] = elements[i]
		}
	}
	return newElements
}
