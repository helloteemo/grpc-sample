package utils

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net"
	"strings"
)

// FilterRepeatIds 过滤重复的id
func FilterRepeatIds(ids []int64) (res []int64) {
	if ids == nil || len(ids) == 0 {
		return
	}
	res = make([]int64, 0, len(ids))
	filterMap := make(map[int64]struct{}, len(ids))

	for _, id := range ids {
		if _, ok := filterMap[id]; ok {
			continue
		}
		filterMap[id] = struct{}{}
		res = append(res, id)
	}
	return
}

// UUID 生成一个uuid，并去掉-，全是小写
func UUID() string {
	value := uuid.New()

	return strings.ToLower(
		strings.Replace(value.String(), "-", "", -1),
	)
}

// GetLocalIp 获取本地IP地址
func GetLocalIp() (string, error) {
	address, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range address {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}

	return "", errors.New("can not find the client ip address")
}

// SplitInt64Array 切分数组
func SplitInt64Array(data []int64, chunkSize int) [][]int64 {
	var result [][]int64
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		result = append(result, data[i:end])
	}
	return result
}
