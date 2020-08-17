package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"os"
	"strconv"
)

func suppressSensitiveFieldDiffs(k, old, new string, d *schema.ResourceData) bool {
	var suppressSensitiveDiff bool
	if old == "" && new != "" {
		return suppressSensitiveDiff
	}
	if suppressSensitiveEnvStr, ok := os.LookupEnv("AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF"); ok {
		if suppressDiff, err := strconv.ParseBool(suppressSensitiveEnvStr); err != nil {
			log.Printf("[ERROR] Invalid value for enviornment variable AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF. Accepted values 1, t, T, true, TRUE, True, 0, f, F, false, FALSE, False. %v", err.Error())
			suppressSensitiveDiff = false
		} else {
			suppressSensitiveDiff = suppressDiff
		}
	}
	return suppressSensitiveDiff
}
