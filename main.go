package main

import (
  "fmt"
)

const (RIGHT = iota
       DOWN
       LEFT
       UP)

func isVisited(spiralsize int, spiral []int, index int, direction int) bool {
  switch direction {
    case RIGHT:
      return spiral[index+1] > 0
    case DOWN:
      return spiral[index+spiralsize] > 0
    case LEFT:
      return spiral[index-1] > 0
    case UP:
      return spiral[index-spiralsize] > 0
  }
  return false
}

func euler28() {
  // Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

  // 21 22 23 24 25
  // 20  7  8  9 10
  // 19  6  1  2 11
  // 18  5  4  3 12
  // 17 16 15 14 13

  // It can be verified that the sum of the numbers on the diagonals is 101.

  // What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
  const spiralsize int = 1001
  spiral := make([]int, spiralsize*spiralsize)

  start := spiralsize / 2
  index := start * spiralsize + start

  // initialization (also the first move to direction)
  count := 1
  spiral[index] = count
  direction := RIGHT

  // direction circle: right, down, left, up, etc...
  for i := 0; i < spiralsize*spiralsize-1; i++ {
    // go to direction
    switch direction {
      case RIGHT:
        index += 1
        // fmt.Println("RIGHT")
      case DOWN:
        index += spiralsize
        // fmt.Println("DOWN")
      case LEFT:
        index -= 1
        // fmt.Println("LEFT")
      case UP:
        index -= spiralsize
        // fmt.Println("UP")
    }

    // increment count
    count += 1
    // set value in spiral
    spiral[index] = count

    // determine new direction
    if direction == RIGHT && !isVisited(spiralsize, spiral, index, DOWN) {
      direction = DOWN
    } else if direction == DOWN && !isVisited(spiralsize, spiral, index, LEFT) {
      direction = LEFT
    } else if direction == LEFT && !isVisited(spiralsize, spiral, index, UP) {
      direction = UP
    } else if direction == UP && !isVisited(spiralsize, spiral, index, RIGHT) {
      direction = RIGHT
    }

    // fmt.Println("new direction:", direction)
  }

  // fmt.Println(spiral)

  diagonalssum := -1 // -1 to correct for the middle '1' otherwise being counted twice
  for i, j := 0, spiralsize*spiralsize-spiralsize; i < spiralsize*spiralsize; {
    ival := spiral[i]
    jval := spiral[j]
    diagonalssum += ival
    diagonalssum += jval
    // fmt.Println("i:", i)
    // fmt.Println("j:", j)
    // fmt.Println("ival:", ival)
    // fmt.Println("jval:", jval)
    i += spiralsize + 1
    j -= spiralsize - 1
  }
  fmt.Println(diagonalssum)
}

func main() {
  euler28()
}
