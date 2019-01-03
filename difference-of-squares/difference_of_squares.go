//Package diffsquares allows you to find the difference between the square of the sum and the sum of the squares of the
// first N natural numbers.
package diffsquares

//SquareOfSum will return the square of the sum of a given number
func SquareOfSum(input int) int {
	sum := (input * (input + 1)) / 2
	return sum * sum
}

//SumOfSquares will return the sum of the squares of the given number
func SumOfSquares(input int) int {
	return (input * (input + 1) * (2 * input + 1)) / 6
}

//Difference will return the difference between the sqaure of sum and the sum of squares
func Difference(input int) int {
	squareOfSum := SquareOfSum(input)
	sumOfSquare := SumOfSquares(input)
	return squareOfSum - sumOfSquare
}