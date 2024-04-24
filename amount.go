package utils

import "math/big"

// FormatAmount 格式化金额字符串
func FormatAmount(amount *big.Float) string {
	return amount.Text('f', 2)
}
