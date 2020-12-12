package lis

type Dyp struct {
	nums     []int
	dp       []int
	longtest int
}

func NewDyp(in []int) *Dyp {
	nums := make([]int, len(in))
	copy(nums, in)
	return &Dyp{
		nums: nums,
	}
}

func (d *Dyp) _Init(mode string) {
	switch mode {
	case "LIS":
		//Both len and cap shall be initialized in case of addtional allocs
		d.dp = make([]int, 1, len(d.nums))
		d.dp[0] = d.nums[0]
	case "LISdynamic":
		d.dp = make([]int, len(d.nums))
	default:
		panic("Invalid LIS method")
	}
}

// LIS binary search
func (d *Dyp) LIS() int {
	d._Init("LIS")
	for _, v := range d.nums {
		if v > d.dp[len(d.dp)-1] {
			d.dp = append(d.dp, v)
		} else if v == d.dp[len(d.dp)-1] {
			continue
		} else {
			d.binaryUpdate(v)
		}
	}
	return len(d.dp)
}

//LISdynamic dynamic process
func (d *Dyp) LISdynamic() int {
	d._Init("LISdynamic")
	for idx := range d.dp {
		d.dp[idx] = 1
	}
	max := 1
	for i := 0; i < len(d.nums); i++ {
		for j := 0; j < i; j++ {
			if d.nums[i] > d.nums[j] && d.dp[j]+1 > d.dp[i] {
				d.dp[i] = d.dp[j] + 1
			}
		}
		if d.dp[i] > max {
			max = d.dp[i]
		}
	}
	return max
}

func (d *Dyp) binaryUpdate(a int) {
	s := d.dp
	for len(s) > 1 {
		mid := len(s) / 2
		if a > s[mid] {
			s = s[mid+1:]
		} else {
			s = s[:mid]
		}
	}
	s[0] = a
}
