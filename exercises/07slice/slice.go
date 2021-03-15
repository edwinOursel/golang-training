package slice

// we will now see how arrays and slices work
// at anytime, to make sure you are working with the right structures,
// you can print their types like this: fmt.Printf("%T\n", myStruct)
// []int indicates a slice while [length]int indicates an array

// first, create an array of 10 integers, 0 to 9 for example, and use it as parameter in a function sumMiddle
// then, use a slice to extract everything except the first and last 3 elements and return the sum of these integers
// you can use the range keyword to loop over the values of a slice

// during your tests while you're at it, try to change the values of one element in one of these slices and print the original array, see what happened ?
// you can also try to print len(slice) and cap(slice) before and after redefining it's upper and lower limit to see how it works
