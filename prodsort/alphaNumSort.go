package prodsort

//AlphaNum AlphaNum
type AlphaNum struct {
	sizeMap map[string]int
}

//Get Get
func (a *AlphaNum) Get() *AlphaNum {
	a.init()
	return a
}

//Sort Sort
func (a *AlphaNum) Sort(s1, s2 string) bool {
	var rtn bool
	i1 := a.sizeMap[s1]
	i2 := a.sizeMap[s2]

	if i1 < i2 {
		rtn = true
	}
	return rtn
}

func (a *AlphaNum) init() {
	a.sizeMap = make(map[string]int)
	a.sizeMap["4XS"] = 1
	a.sizeMap["3XS"] = 2
	a.sizeMap["2XS"] = 3
	a.sizeMap["XS"] = 4
	a.sizeMap["S"] = 5
	a.sizeMap["M"] = 6
	a.sizeMap["L"] = 7
	a.sizeMap["XL"] = 8
	a.sizeMap["2XL"] = 9
	a.sizeMap["XXL"] = 9
	a.sizeMap["3XL"] = 10
	a.sizeMap["XXXL"] = 10
	a.sizeMap["4XL"] = 11
	a.sizeMap["XXXXL"] = 11
	a.sizeMap["5XL"] = 12
	a.sizeMap["XXXXXL"] = 12
	a.sizeMap["6XL"] = 14
	a.sizeMap["XXXXXXL"] = 14
	a.sizeMap["7XL"] = 15
	a.sizeMap["8XL"] = 16
	a.sizeMap["9XL"] = 17
	a.sizeMap["10XL"] = 18
	a.sizeMap["11XL"] = 19
	a.sizeMap["12XL"] = 20
	a.sizeMap["13XL"] = 21
	a.sizeMap["14XL"] = 22
	a.sizeMap["15XL"] = 23
	//tall
	a.sizeMap["ST"] = 30
	a.sizeMap["MT"] = 31
	a.sizeMap["LT"] = 32
	a.sizeMap["XLT"] = 33
	a.sizeMap["2XLT"] = 34
	a.sizeMap["3XLT"] = 35
	a.sizeMap["4XLT"] = 36
	a.sizeMap["5XLT"] = 37
	a.sizeMap["6XLT"] = 38
	a.sizeMap["7XLT"] = 39
	a.sizeMap["8XLT"] = 40
	a.sizeMap["9XLT"] = 41
	a.sizeMap["10XLT"] = 42
	a.sizeMap["11XLT"] = 43
	a.sizeMap["12XLT"] = 44
	a.sizeMap["13XLT"] = 45
	a.sizeMap["14XLT"] = 46

	//shoe sizes M L and half sizes start at 50

	a.sizeMap["5"] = 50
	a.sizeMap["5M"] = 51
	a.sizeMap["5W"] = 52

	a.sizeMap["5.5"] = 53
	a.sizeMap["5.5M"] = 54
	a.sizeMap["5.5W"] = 55

	a.sizeMap["6"] = 56
	a.sizeMap["6M"] = 57
	a.sizeMap["6W"] = 58

	a.sizeMap["6.5"] = 59
	a.sizeMap["6.5M"] = 60
	a.sizeMap["6.5W"] = 61

	a.sizeMap["7"] = 62
	a.sizeMap["7M"] = 63
	a.sizeMap["7W"] = 64

	a.sizeMap["7.5"] = 65
	a.sizeMap["7.5M"] = 66
	a.sizeMap["7.5W"] = 67

	a.sizeMap["8"] = 68
	a.sizeMap["8M"] = 69
	a.sizeMap["8W"] = 70

	a.sizeMap["8.5"] = 71
	a.sizeMap["8.5M"] = 72
	a.sizeMap["8.5W"] = 73

	a.sizeMap["9"] = 74
	a.sizeMap["9M"] = 75
	a.sizeMap["9W"] = 76

	a.sizeMap["9.5"] = 77
	a.sizeMap["9.5M"] = 78
	a.sizeMap["9.5W"] = 79

	a.sizeMap["10"] = 80
	a.sizeMap["10M"] = 81
	a.sizeMap["10W"] = 82

	a.sizeMap["10.5"] = 83
	a.sizeMap["10.5M"] = 84
	a.sizeMap["10.5W"] = 85

	a.sizeMap["11"] = 86
	a.sizeMap["11M"] = 87
	a.sizeMap["11W"] = 88

	a.sizeMap["11.5"] = 89
	a.sizeMap["11.5M"] = 90
	a.sizeMap["11.5W"] = 91

	a.sizeMap["12"] = 92
	a.sizeMap["12M"] = 93
	a.sizeMap["12W"] = 94

	a.sizeMap["12.5"] = 95
	a.sizeMap["12.5M"] = 96
	a.sizeMap["12.5W"] = 97

	a.sizeMap["13"] = 98
	a.sizeMap["13M"] = 99
	a.sizeMap["13W"] = 100

	a.sizeMap["13.5"] = 101
	a.sizeMap["13.5M"] = 102
	a.sizeMap["13.5W"] = 103

	a.sizeMap["14"] = 104
	a.sizeMap["14M"] = 105
	a.sizeMap["14W"] = 106

	a.sizeMap["14.5"] = 107
	a.sizeMap["14.5M"] = 108
	a.sizeMap["14.5W"] = 109

	a.sizeMap["15"] = 110
	a.sizeMap["15M"] = 111
	a.sizeMap["15W"] = 112

	a.sizeMap["15.5"] = 113
	a.sizeMap["15.5M"] = 114
	a.sizeMap["15.5W"] = 115

	a.sizeMap["16"] = 116
	a.sizeMap["16M"] = 117
	a.sizeMap["16W"] = 118

	a.sizeMap["16.5"] = 119
	a.sizeMap["16.5M"] = 120
	a.sizeMap["16.5W"] = 121

	a.sizeMap["17"] = 122
	a.sizeMap["17M"] = 123
	a.sizeMap["17W"] = 124

	a.sizeMap["17.5"] = 125
	a.sizeMap["17.5M"] = 126
	a.sizeMap["17.5W"] = 127

}
