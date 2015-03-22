 From      | `b []byte`     | `i []int`      | `r []rune`     | `s string`     | `f flt32`      | `i int`
-----------|----------------|----------------|----------------|----------------|----------------|---------
     **To**|                |                |                |                |                |
  `[]byte` |      ·         |                |                | `[]byte(s)`    |                |
  `[]int`  |                |       ·        |                | `[]int(s)`     |                |
  `[]rune` |                |                |                | `[]rune(s)`    |                |
  `string` | `string(b)`    | `string(i)`    | `string(r)`    |       ·        |                |
 `ftl32`   |                |                |                |                |        ·       | `flt32(i)`
     `int` |                |                |                |                | `int(f)`       |     ·
Table: Valid conversions, `float64` works the same as `float32`. Note that
float32 has been abbreviated to flt32 in this table to make it fit on the page.
