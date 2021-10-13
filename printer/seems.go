package printer

import (
	"strconv"
	"strings"

	"github.com/SignorMercurio/limner/color"
)

func seemsNegative(status string) bool {
	negativeKeywords := []string{
		"fail",
		"backoff",
		"exceed",
		"not",
		"err",
		"invalid",
		"unable",
		"unhealthy",
		"unknown",
		"unavailable",
		"evict",
		"bad",
		"timeout",
		"panic",
		"fatal",
	}

	for _, v := range negativeKeywords {
		if strings.Contains(status, v) {
			return true
		}
	}
	return false
}

func seemsWarning(status string) bool {
	return strings.Contains(status, "ing")
}

func seemsPositive(status string) bool {
	positiveKeywords := []string{
		"ok",
		"ted",
		"led",
		"ged",
		"zed",
		"success",
		"succeed",
		"ready",
		"normal",
		"healthy",
		"running",
		"done",
		"available",
	}

	for _, v := range positiveKeywords {
		if strings.Contains(status, v) {
			return true
		}
	}
	return false
}

func seemsReadyStatus(status string) (color.Color, bool) {
	if strings.Count(status, "/") == 1 {
		ready := strings.Split(status, "/")
		if ready[0] == ready[1] {
			return color.Green, true
		} else {
			_, e1 := strconv.Atoi(ready[0])
			_, e2 := strconv.Atoi(ready[1])
			if e1 == nil && e2 == nil {
				return color.Yellow, true
			}
		}
	}
	return 0, false
}
