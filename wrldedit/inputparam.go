package wrldedit

import (
	"errors"
	"strconv"
	"strings"
)

type InputParam struct {
	param map[string]string
}

type ParameterWrongError error

func ReadTileInfo(input string) (InputParam, error) {
	info := make(map[string]string)

	split := strings.Split(input, " ")
	info["Name"] = split[0]

	if len(split) < 2 || split[1] == "" {
		return InputParam{info}, nil
	}

	params := split[1]
	for _, param := range strings.Split(params, ",") {
		parsplit := strings.Split(param, ":")

		if len(parsplit) != 2 {
			return InputParam{info}, errors.New("Failed Reading Parameters")
		}
		pname := parsplit[0]
		parg := parsplit[1]

		info[pname] = parg
	}

	return InputParam{info}, nil
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

func (ip InputParam) GetInt(key string) (int, bool) {
	str, avab := ip.param[key]

	if !avab {
		return 0, false
	}

	nt, err := strconv.Atoi(str)

	if err != nil {
		return 0, false
	}

	return nt, true
}

func (ip InputParam) GetIntElse(key string, els int) int {
	nt, avab := ip.GetInt(key)

	if avab {
		return nt
	} else {
		return els
	}
}
