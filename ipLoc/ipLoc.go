package ipLoc

import (
	"errors"
	"strconv"
	"strings"

	"github.com/nightin-gale/ipLoc/data"
)

// A datatype for storing ip infomations
type IpDetail struct {
	Ip          string
	IpInt       uint64
	CountryCode string
	Region      string
	City        string
}

// Convert an ip in form of string like "1.2.3.10" to uint64 representation
// "a.b.c.d" -> a*256^3 + b*256^2 + c*256 + d
func IpToUint64(ip string) (uint64, error) {
	spl := strings.Split(ip, ".")

	if len(spl) != 4 {
		return 0, errors.New("Invalid IP address")
	}
	var power2 []uint64 = []uint64{1 << 24, 1 << 16, 1 << 8, 1}
	var num uint64 = 0
	for i := 0; i < 4; i++ {
		tmp, err := strconv.ParseUint(spl[i], 10, 64)
		if err != nil {
			return 0, err
		}
		num += tmp * power2[i]
	}
	return num, nil
}

func Uint64ToIp(intIp uint64) (string, error) {
	var power2 []uint64 = []uint64{1 << 24, 1 << 16, 1 << 8, 1}
	var ip string = ""
	for i := 0; i < 4; i++ {
		var tmp uint64 = intIp / power2[i]
		intIp = intIp % power2[i]
		u := strconv.FormatUint(tmp, 10)
		ip += u
		if i != 3 {
			ip += "."
		}
	}
	return ip, nil
}

func binarySearch(array []uint64, target uint64) (int, error) {
	var lower, upper = 0, len(array) - 1
	var preLower, preUpper = 0, len(array) - 1

	for array[lower] > target || array[lower+1] <= target {
		if array[lower+1] == target {
			return lower + 1, nil
		}
		midpoint := (lower + upper) / 2
		if array[midpoint] < target {
			lower = midpoint
		} else {
			upper = midpoint
		}

		// Error handling
		if lower == preLower && upper == preUpper {
			return lower, errors.New("Unable to find target")
		}
		preLower = lower
		preUpper = upper
	}

	return lower, nil
}

// Look up information about an ipv4 IP address
// The data is stored in two arrays: Ipv4Ip and Ipv4Data. Ipv4Ip is a strictly increasing array of uint64, recording int representation of the Ip address. An Ip address, ip, that is Ipv4Ip[i] <= ip <= Ipv4Ip[i+1], has its data stored at Ipv4Data[i].
// These two arrays are stored in "../data/data.go", generated from the data in the file "../script/ipv4location.csv" by the script "../script/processData.go
func IpLoc(ip string) (IpDetail, error) {
	target, err := IpToUint64(ip)
	if err != nil {
		return IpDetail{}, err
	}

	index, err := binarySearch(data.Ipv4Ip, target)
	if err != nil {
		return IpDetail{}, err
	}
	return IpDetail{
			Ip:          ip,
			IpInt:       target,
			CountryCode: string(data.Ipv4Data[index].Country[:]),
			Region:      data.Ipv4Data[index].Region,
			City:        data.Ipv4Data[index].City,
		},
		nil
}
