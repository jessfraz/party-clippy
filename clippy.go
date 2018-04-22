package main

const (
	clippyBlurb = `
 _________________________________
/ It looks like you're building a \
\ microservice.                   /
 ---------------------------------
 \
  \
`

	clippyBody0 = `     __
    /  \
    |  |
    @  @
    |  |
    || |/
    || ||
    |\_/|
    \___/
`
	clippyBody1 = `     __
    /  \
    |  |
     @  @
    |  |
    || |/
    || ||
    |\_/|
    \___/
`
	clippyBody2 = `     __
    /  \
    |  |
    @  @
    |  |
    || |/
    || ||
    |\_/|
    \___/
`
	clippyBody3 = `     __
    /  \
    |  |
    @  @
    |  |
    || | /
    || ||
    |\_/|
    \___/
`
	clippyBody4 = `     __
    /  \
    |  |
    @  @
    |  |
    || |  /
    || ||
    |\_/|
    \___/
`
	clippyBody5 = `     __
    /  \
    |  |
    @  @
    |  |
    || | /
    || ||
    |\_/|
    \___/
`
)

var (
	clippyBodies = map[int]string{
		0: clippyBody0,
		1: clippyBody1,
		2: clippyBody2,
		3: clippyBody3,
		4: clippyBody4,
		5: clippyBody5,
	}
)

func wiggle(i int) string {
	return clippyBodies[i]
}
