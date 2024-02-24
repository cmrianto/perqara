package helper

import (
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetString(in *string) string {
	res := ""

	if in != nil {
		res = *in
	}

	return res
}

func GetPointerStringToPointer(in *string) *string {
	if in != nil {
		return in
	}

	return nil
}

func GetStringNotEmpty(in string) *string {
	if len(in) > 0 {
		return &in
	}

	return nil
}

func GetInt32(in *int32) int32 {
	var res int32
	res = 0

	if in != nil {
		res = *in
	}

	return res
}

func GetInt64(in *int64) int64 {
	var res int64
	res = 0

	if in != nil {
		res = *in
	}

	return res
}

func GetPointerInt64ToPointer(in *int64) *int64 {
	if in != nil {
		return in
	}

	return nil
}

func GetInt64NotEmpty(in int64) *int64 {
	if in > 0 {
		return &in
	}

	return nil
}

func GetInt(in *int) int {
	res := 0

	if in != nil {
		res = *in
	}

	return res
}

func GetFloat64(in *float64) float64 {
	res := float64(0)
	if in != nil {
		res = *in
	}

	return res
}

func ToString(in *string) string {
	if in != nil {
		return *in
	}
	return ""
}

func ToInt32(in *int32) int32 {
	if in != nil {
		return *in
	}
	return 0
}

func ToFloat64(in *float64) float64 {
	if in != nil {
		return *in
	}
	return 0
}

func ToBool(in *bool) bool {
	if in != nil {
		return *in
	}
	return false
}

func MakeArray(in *string, sep string) []string {
	if ToString(in) != "" {
		return strings.Split(*in, sep)
	}
	return nil
}

func MakeArrayInt32(in *string, sep string) []int32 {
	if ToString(in) != "" {
		ret := []int32{}
		splitted := strings.Split(*in, sep)
		for _, i := range splitted {
			ii, _ := strconv.Atoi(i)
			ret = append(ret, int32(ii))
		}
		return ret
	}
	return nil
}

func MakeArrayInt64(in *string, sep string) []int64 {
	if ToString(in) != "" {
		ret := make([]int64, 0)
		splitted := strings.Split(*in, sep)
		for _, i := range splitted {
			ii, _ := strconv.ParseInt(strings.TrimSpace(i), 10, 64)
			ret = append(ret, ii)
		}
		return ret
	}
	return nil
}

func ToFloat32(in *float32) float32 {
	if in != nil {
		return *in
	}
	return 0
}

func GetStringDefault(in *string, def string) string {
	if ToString(in) != "" {
		return *in
	}
	return def
}

func ToFloat64Pointer(in float64) *float64 {
	return &in
}

func ToStringPointer(in string) *string {
	return &in
}

func ToBoolPointer(in bool) *bool {
	return &in
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
