package versionutil

import (
	"regexp"
	"strconv"
)

func TrimVersionSuffix(ver string) string {
	re := regexp.MustCompile(`^([0-9]+(\.[0-9]+)*)`)
	match := re.FindStringSubmatch(ver)
	if len(match) > 1 {
		return match[1]
	}
	return ver // 如果没有匹配则原样返回
}

func VersionInRange(curVer, minVer, maxVer string) bool {
	vstr := TrimVersionSuffix(curVer)
	v := parseVersion(vstr)
	vmin := parseVersion(minVer)
	vmax := parseVersion(maxVer)
	llen := maxInt(len(v), len(vmin), len(vmax))
	v = padVersion(v, llen)
	vmin = padVersion(vmin, llen)
	vmax = padVersion(vmax, llen)

	// Compare with minVer
	if compareVersion(v, vmin) < 0 {
		return false
	}
	// Compare with maxVer
	if compareVersion(v, vmax) > 0 {
		return false
	}
	return true
}

func parseVersion(ver string) []int {
	re := regexp.MustCompile(`[0-9]+`)
	parts := re.FindAllString(ver, -1)
	res := make([]int, len(parts))
	for i, p := range parts {
		res[i], _ = strconv.Atoi(p)
	}
	return res
}

func padVersion(v []int, n int) []int {
	for len(v) < n {
		v = append(v, 0)
	}
	return v
}

func compareVersion(a, b []int) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}
	return 0
}

func maxInt(a, b, c int) int {
	m := a
	if b > m {
		m = b
	}
	if c > m {
		m = c
	}
	return m
}
