package combinnations

// Combinnations 组合描述
type Combinnations struct {
	n    int
	k    int
	comb []int
}

// NewConmbinnations ...
func NewConmbinnations(n, k int) *Combinnations {
	c := &Combinnations{}
	if k > n || k <= 0 {
		return nil
	}
	c.n = n
	c.k = k
	c.comb = make([]int, k)
	for i := 0; i < k; i++ {
		c.comb[i] = i
	}
	return c
}

// Next 生成下一个组合
func (c *Combinnations) Next() bool {
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

// Comb 取出组合
func (c *Combinnations) Comb() []int {
	return c.comb
}
