package middle

//impl Solution {
//	pub fn my_pow(x: f64, n: i32) -> f64 {
//		if n < 0 {
//			return 1.0 / Self::square(x, -n);
//		} else {
//			return Self::square(x, n);
//		}
//	}
//
//	pub fn square(result: f64, tmp: i32) -> f64 {
//		if tmp == 0 {
//			return 1.0;
//		}
//
//		if tmp % 2 == 1 {
//			let y = Self::square(result, tmp / 2);
//			return result * y * y;
//		} else {
//			let y = Self::square(result, tmp / 2);
//			return y * y;
//		}
//	}
//}

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, n)
	} else {
		return pow(x, n)
	}

}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		y := pow(x, n/2)
		return y * y * x
	} else {
		y := pow(x, n/2)
		return y * y
	}
}
