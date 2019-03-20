package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type surfaceQuery struct {
	width     int
	height    int
	xyscale   float64
	zscale    float64
	topRGB    string
	bottomRGB string
}

func parseQuery(qstr string) *surfaceQuery {
	query := surfaceQuery{width, height, xyscale, zscale, "red", "blue"}
	for _, namedValueStr := range strings.Split(qstr, ",") {
		fillQuery(namedValueStr, &query)
	}
	return &query
}

func getColor(z float64, query *surfaceQuery) string {
	if z > 0 {
		return toStrColor(z, query.topRGB)
	} else {
		return toStrColor(z, query.bottomRGB)
	}
}

func fillQuery(namedValueStr string, query *surfaceQuery) {
	namedValue := strings.Split(namedValueStr, "=")
	if len(namedValue) != 2 {
		fmt.Fprintf(os.Stderr, "[%s] is not supported\n", namedValueStr)
		return
	}
	if namedValue[0] == "width" {
		query.width = parseInt(namedValue[1], query.width)
		query.xyscale = float64(query.width) / 2 / xyrange
	} else if namedValue[0] == "height" {
		query.height = parseInt(namedValue[1], query.height)
		query.zscale = float64(query.height) * 0.4
	} else if namedValue[0] == "topRGB" && isRGB(namedValue[1]) {
		query.topRGB = namedValue[1]
	} else if namedValue[0] == "bottomRGB" && isRGB(namedValue[1]) {
		query.bottomRGB = namedValue[1]
	} else {
		fmt.Fprintf(os.Stderr, "[%s] is not supported\n", namedValueStr)
	}
}

func parseInt(s string, defaultValue int) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "strconv.ParseInt: %v\n", err)
		return defaultValue
	}
	return int(i)
}

func isRGB(value string) bool {
	return ((value == "red") || (value == "green") || (value == "blue"))
}

func toStrColor(z float64, rgb string) string {
	color := int(0xff * (math.Abs(z)))
	if color > 0xff {
		color = 0xff
	}
	strColor := fmt.Sprintf("%02x", color)

	if rgb == "red" {
		return ("#" + strColor + "0000")
	} else if rgb == "green" {
		return ("#00" + strColor + "00")
	} else if rgb == "blue" {
		return ("#0000" + strColor)
	}
	return "#000000"
}
