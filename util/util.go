package util

import (
	"marvin/GraphEng/GE"
)

//IndexOf()
func IndexOf(slice []*GE.Button, btn *GE.Button) int {
	for x, val := range slice {
		if val == btn {
			return x
		}
	}

	return -1
}
