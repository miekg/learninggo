 From      | `b []byte`     | `i []int`      | `r []rune`     | `s string`     | `f float32`    | `i int`
-----------|----------------|----------------|----------------|----------------|----------------|-------------
     **To**|                |                |                |                |                |
  `[]byte` |      ·         |                |                | `[]byte(s)`    |                |
  `[]int`  |                |       ·        |                | `[]int(s)`     |                |
  `[]rune` |                |                |                | `[]rune(s)`    |                |
  `string` | `string(b)`    | `string(i)`    | `string(r)`    |       ·        |                |
 `float32` |                |                |                |                |        ·       | `float32(i)`
     `int` |                |                |                |                | `int(f)`       |     ·
Table: Valid conversions, `float64` works the same as `float32`.
