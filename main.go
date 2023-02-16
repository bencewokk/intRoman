type number struct {
	integer int
	roman   string
}

func intToRoman(num int) string {

	I := number{integer: 1, roman: "I"}
	II := number{integer: 2, roman: "II"}
	III := number{integer: 3, roman: "III"}
	IV := number{integer: 4, roman: "IV"}
	V := number{integer: 5, roman: "V"}
	IX := number{integer: 9, roman: "IX"}
	X := number{integer: 10, roman: "X"}
	XX := number{integer: 20, roman: "XX"}
	XXX := number{integer: 30, roman: "XXX"}
	XL := number{integer: 40, roman: "XL"}
	L := number{integer: 50, roman: "L"}
	XC := number{integer: 90, roman: "XC"}
	C := number{integer: 100, roman: "C"}
	CC := number{integer: 200, roman: "CC"}
	CCC := number{integer: 300, roman: "CCC"}
	CD := number{integer: 400, roman: "CD"}
	D := number{integer: 500, roman: "D"}
	CM := number{integer: 900, roman: "CM"}
	M := number{integer: 1000, roman: "M"}
	MM := number{integer: 2000, roman: "MM"}
	MMM := number{integer: 3000, roman: "MMM"}

	var output string
	nums := []number{I, II, III, IV, V, X, IX, X, XX, XXX, XL, L, XC, C, CC, CCC, CD, D, CM, M, MM, MMM}
	var numsRn int
	u := len(nums) - 1

	for {
		//fmt.Println(u, numsRn)
		//fmt.Println(u, nums[u].integer)
		numsRn = num - nums[u].integer
		if numsRn > 0 || numsRn == 0 {
			num -= nums[u].integer
			output = output + nums[u].roman
		}

		u--

		if numsRn == 0 {
			break
		}
	}

	//fmt.Println(output)

	return output
}
