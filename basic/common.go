package basic

import (
	"github.com/wesovilabs/koazee-comparison/util"
)

var numbers5000 = util.ArrayOfInt(0, 2000, 5000)
var strings5000 = util.ArrayOfString(1, 10, 5000)
var sum = func(acc, elem int) int { return acc + elem }
