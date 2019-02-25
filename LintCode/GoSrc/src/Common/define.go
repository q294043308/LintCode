package Common

import (
	"math"
)

var MAXWINTNUM = math.Pow(2, 61) - 1
var MAXINTNUM int = int(math.Pow(2, 31) - 1)
var MININTNUM int = int(-(math.Pow(2, 31)) + 1)

const CHARNUM = 128
const ENGLISH_CHAR_NUM = (26 * 2)
const BIG_ENGLISH_CHAR_NUM = 26
const SMALL_ENGLISH_CHAR_NUM = 26
const MIN_LEN = 0x7fffffff
const CIRCLE_PI = 3.14
const THIRTY = 30

type Point struct {
	X int
	Y int
}
