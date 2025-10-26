package shortener

const (
	charset      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	MaxUrlLength = 6
)

func ToBase62(num int64) string {
	if num == 0 {
		return string(charset[0])
	}

	result := ""

	for num > 0 {
		remainder := num % 62
		result = string(charset[remainder]) + result
		num = num / 62
	}

	if len(result) < MaxUrlLength {
		for len(result) < MaxUrlLength {
			result = "0" + result
		}
	}

	return result
}

func FromBase62(str string) int64 {
	var result int64 = 0
	for i := 0; i < len(str); i++ {
		char := str[i]
		var value int64
		if char >= '0' && char <= '9' {
			value = int64(char - '0')
		} else if char >= 'A' && char <= 'Z' {
			value = int64(char - 'A' + 10)
		} else if char >= 'a' && char <= 'z' {
			value = int64(char - 'a' + 36)
		}
		result = result*62 + value
	}

	return result
}
