package math

func myPow(x float64, n int) float64 {
	var quickM func(x float64, n int) float64
	quickM = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		y := quickM(x, n/2)
		if n%2 == 0 {
			return y * y
		} else {
			return y * y * x
		}
	}

	if n > 0 {
		return quickM(x, n)
	} else {
		return 1.0 / quickM(x, -n)
	}
}
