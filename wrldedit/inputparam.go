package wrldedit

import (
	"strconv"
	"strings"
)

type InputParam struct {
	param map[string]string
}

func ReadTileInfo(input string) InputParam {
	info := make(map[string]string)

	split := strings.Split(input, " ")
	info["Name"] = split[0]

	if len(split) < 2 {
		return InputParam{info}
	}

	params := strings.Split(input, " ")[1]
	for _, param := range strings.Split(params, ",") {
		pname := strings.Split(param, ":")[0]
		parg := strings.Split(param, ":")[1]

		info[pname] = parg
	}

	return InputParam{info}
}

func (ip InputParam) GetString(key string) (string, bool) {
	str, avab := ip.param[key]
	return str, avab
}

func (ip InputParam) GetStringElse(key, els string) string {
	str, avab := ip.param[key]

	if avab {
		return str
	} else {
		return els
	}
}

func (ip InputParam) GetFloat64(key string) (float64, bool) {
	str, avab := ip.param[key]

	if !avab {
		return 0, false
	}

	flt, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return 0, false
	}

	return flt, true
}

func (ip InputParam) GetFloat64Else(key string, els float64) float64 {
	flt, avab := ip.GetFloat64(key)

	if avab {
		return flt
	} else {
		return els
	}
}