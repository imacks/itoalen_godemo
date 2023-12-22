package itoalentest

import (
	"math"
	"strconv"
)

func ItoalenWithItoa(n int) int {
	return len(strconv.Itoa(n))
}

func ItoalenWithLog(n int) int {
	if n == 0 {
		return 1
	}

	var prefixlen int
	if n < 0 {
		prefixlen = 1
		n = n*-1
	}
	return int(math.Log10(float64(n))) + 1 + prefixlen
}

func ItoalenWithMul(n int) int {
	if n == 0 {
		return 1
	}

	var prefixlen int
	if n < 0 {
		prefixlen = 1
		n = n*-1
	}
	if n >= 1000000000000000000 {
		return 19
	}

	var nlen int
	for nn := 1; nn <= n; nn*=10 {
		nlen++
	}
	return nlen + prefixlen
}

func ItoalenWithMul2(n int) int {
	if n == 0 {
		return 1
	}

	var prefixlen int
	if n < 0 {
		prefixlen = 1
		n = n*-1
	}
	if n >= 1000000000000000000 {
		return 19
	}

	var nlen int
	nn := 1
	for nn <= n {
		nlen++
		nn = nn << 3 + nn << 1
	}
	return nlen + prefixlen
}

func ItoalenWithDiv2(n int) int {
	var prefixlen int
	if n < 0 {
		prefixlen = 1
		n = n*-1
	}

	nlen := 1

	if n >= 10000000000000000 {
		nlen += 16
		n /= 10000000000000000
	}

	if n >= 100000000 {
		nlen += 8
		n /= 100000000
	}

	if n >= 100000000 {
		nlen += 8
		n /= 100000000
	}

	if n >= 10000 {
		nlen += 4
		n /= 10000
	}

	if n >= 100 {
		nlen += 2
		n /= 100
	}

	if n >= 10 {
		nlen++
	}

	return nlen + prefixlen
}

func ItoalenWithBranch(n int) int {
	var prefixlen int
	if n < 0 {
		prefixlen = 1
		n = n*-1
	}

	if n < 10000000000 {
		if n < 100000 {
			if n < 1000 {
				if n < 100 {
					if n < 10 {
						return 1 + prefixlen
					} else {
						return 2 + prefixlen
					}
				} else {
					return 3 + prefixlen
				}
			} else {
				if n < 10000 {
					return 4 + prefixlen
				} else {
					return 5 + prefixlen
				}
			}
		} else {
			if n < 100000000 {
				if n < 10000000 {
					if n < 1000000 {
						return 6 + prefixlen
					} else {
						return 7 + prefixlen
					}
				} else {
					return 8 + prefixlen
				}
			} else {
				if n < 1000000000 {
					return 9 + prefixlen
				} else {
					return 10 + prefixlen
				}
			}
		}
	} else {
		if n < 1000000000000000 {
			if n < 10000000000000 {
				if n < 1000000000000 {
					if n < 100000000000 {
						return 11 + prefixlen
					} else {
						return 12 + prefixlen
					}
				} else {
					return 13 + prefixlen
				}
			} else {
				if n < 100000000000000 {
					return 14 + prefixlen
				} else {
					return 15 + prefixlen
				}
			}
		} else {
			if n < 1000000000000000000 {
				if n < 100000000000000000 {
					if n < 10000000000000000 {
						return 16 + prefixlen
					} else {
						return 17 + prefixlen
					}
				} else {
					return 18 + prefixlen
				}
			} else {
				return 19 + prefixlen
			}
		}
	}
}