package tools

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

//时间格式转时间戳
func Str2Timestamp(timestr string) int64 {
	timeLayout := "2006-01-02T15:04:05.999"                      //转化所需模板
	loc, _ := time.LoadLocation("Local")                         //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, timestr, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                         //转化为时间戳 类型是int64
	return sr
}

//截取字符串-将字符串前面的0去掉
func SubByZero(str string) string {
	arr := []byte(str)
	for i, str1 := range arr {
		if str1 != 48 {
			str = string(arr[i:])
			break
		}
	}
	return str
}

//参数解析
func ValiedParam(params []string) bool {
	for _, param := range params {
		if param == "" {
			return false
		}
	}
	return true
}

//精度处理
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

//浮点数精度处理转换整数
func Float2Int(f float64, decimal int) uint64 {
	value := ChangeNumber(f, 8)
	split := strings.Split(value, ".")
	if len(split) == 2 {
		decimal -= len(split[1])
	}
	value = strings.ReplaceAll(value, ".", "")
	for i := 0; i < decimal; i++ {
		value += "0"
	}
	i, _ := strconv.ParseInt(value, 10, 64)
	return uint64(i)
}

//保留小数位
func ChangeNumber(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

//将字符串float进行精度处理
func Float2String(value string, decimal int) string {
	split := strings.Split(value, ".")
	if len(split) > 1 {
		decimal -= len(split[1])
		value = strings.ReplaceAll(value, ".", "")
	}
	for decimal > 0 {
		value += "0"
		decimal -= 1
	}
	return SubByZero(value)
}

//精度处理
func ChangeInt2Float(f string, m int) string {
	i := len(f)
	if i <= m {
		for i < m {
			f = "0" + f
			i++
		}
		return "0." + f
	} else {
		sliceTep := make([]string, i)
		for index, value := range f {
			if index == i-m {
				sliceTep = append(sliceTep, ".")
			}
			sliceTep = append(sliceTep, string(value))
		}
		return strings.Join(sliceTep, "")
	}
}

//精度处理
func ChangeFloat2Float(f string, m int) string {
	if strings.Contains(f, ".") {
		index := strings.Index(f, ".")
		sub := f[index+1 : len(f)]
		if len(sub) > m {
			f = f[0 : index+m+1]
		} else {
			for m > len(sub) {
				f += "0"
				m--
			}
		}
	} else {
		f += "."
		for m > 0 {
			f += "0"
			m--
		}
	}
	return f
}

func HexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex[2:], 16)
	return n
}

//小数点位数校验
func CheckFloat(value string, decimal int) bool {
	values := strings.Split(value, ".")
	if len(values) == 2 {
		if len(values[1]) > decimal {
			return false
		}
	}
	return true
}

func GetOrderid() string {
	var t = time.Now().Unix()
	var s = time.Unix(t, 0).Format("20060102150405") + fmt.Sprint(time.Now().UnixNano())
	return s
}

//结构体转为map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

//Str2Str 精度处理
func Str2Str(f string, m int) string {
	i := len(f)
	if i <= m {
		for i < m {
			f = "0" + f
			i++
		}
		return "0." + f
	} else {
		sliceTep := make([]string, i)
		for index, value := range f {
			if index == i-m {
				sliceTep = append(sliceTep, ".")
			}
			sliceTep = append(sliceTep, string(value))
		}
		return strings.Join(sliceTep, "")
	}
}

//String2Hex 将字符串转十六进制
func String2Hex(value string) string {
	int := new(big.Int)
	int.SetString(value, 0)
	return common.Bytes2Hex(int.Bytes())
}
