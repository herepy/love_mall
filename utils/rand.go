/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/10/13 16:29
 */

package utils

import "math/rand"

func RandInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandStr(len int) string {
	b := make([]byte, len)
	for i := 0; i < len; i++ {
		r := RandInt(65, 122)
		if r >= 91 && r <= 96 {
			r = 97
		}
		b[i] = byte(r)
	}

	return string(b)
}
