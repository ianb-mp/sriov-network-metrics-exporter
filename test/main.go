package main

import (
	"fmt"
	"regexp"
)

const (
	VersionRegexpRaw string = `v?([0-9]+(\.[0-9]+)*?)(-([0-9]+[0-9A-Za-z\-~]*(\.[0-9A-Za-z\-~]+)*)|(-?([A-Za-z\-~]+[0-9A-Za-z\-~]*(\.[0-9A-Za-z\-~]+)*)))?(\+([0-9A-Za-z\-~]+(\.[0-9A-Za-z\-~]+)*))??`
)

func main() {

	versionRegexp := regexp.MustCompile("^" + VersionRegexpRaw + "$")

	//v := "5.14.0-362.8.1.el9_3.x86_64"
	v := "5.14.0-362.8.1.el9.3.x8664"
	matches := versionRegexp.FindStringSubmatch(v)
	if matches == nil {
		fmt.Printf("Malformed version: %s\n", v)
	}
	fmt.Printf("matches %v\n", matches)
}
