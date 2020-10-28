/**
** @创建时间: 2020/10/28 10:26 上午
** @作者　　: return
** @描述　　:
 */
package conf

import (
	"fmt"
	"math/rand"
	"testing"
)

// 生成随机的盐字符串
func TestGenerate(t *testing.T) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 18)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	fmt.Println(string(b))
}