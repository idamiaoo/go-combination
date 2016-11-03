package combinnations

type combinnations struct {
	n    int
	k    int
	comb []int
}

//NewConmbinnations ,,,
func NewConmbinnations() *combinnations {
	return &combinnations{}
}

func (c *combinnations) Init(n, k int) bool {
	if k > n || k <= 0 {
		return false
	}
	c.n = n
	c.k = k
	c.comb = make([]int, k)
	for i := 0; i < k; i++ {
		c.comb[i] = i
	}
	return true
}

func (c *combinnations) Next() bool {
	ii := c.k - 1
	for ii >= 0 && c.comb[ii]+c.k+1 > c.n+ii {
		ii--
	}
	if ii < 0 {
		return false
	}
	c.comb[ii]++
	ii++
	for ii < c.k {
		c.comb[ii] = c.comb[ii-1] + 1
		ii++
	}
	return true
}

func (c *combinnations) Comb() []int {
	return c.comb
}
