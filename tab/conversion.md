    *From* | `b []byte`     | `i []int`      | `r []rune`     | `s string`     | `f flt32`      | `i int`
-----------|----------------|----------------|----------------|----------------|----------------|
      *To* |                |                |                |                |                |
  `[]byte` | $$\texttimes$$ |                |                | `[]byte(s)`    |                |
  `[]int`  |                | $$\texttimes$$ |                | `[]int(s)`     |                |
  `[]rune` |                |                | $$\texttimes$$ | `[]rune(s)`    |                |
  `string` | `string(b)`    | `string(i)`    | `string(r)`    | $$\texttimes$$ |                |
 `ftl32`   |                |                |                |                | $$\texttimes$$ | `flt32(i)`
     `int` |                |                |                |                | `int(f)`       | $$\texttimes$$
Figure: Valid conversions, `float64` works the same as `float32`. Note that
float32 has been abbreviated to flt32 in this table to make it fit on the page.
