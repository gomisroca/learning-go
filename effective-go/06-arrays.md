## Arrays

Arrays are primaly used as building blocks for slices.

In Go, arrays are values: assigning an array to another copies all elements. If we pass an array to a function, the function will get a copy of the array, not a pointer to it. The size of the array is part of its type, [10]int and [20]int are different types.

In general, we should avoid using arrays and use slices instead.
