/***************************************************************************
 * ========================================================================
 * Copyright 2022 VMware, Inc.  All rights reserved. VMware Confidential
 * ========================================================================
 */

package avi

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func suppressSensitiveFieldDiffs(k, old, new string, d *schema.ResourceData) bool {
	var suppressSensitiveDiff bool
	if old == "" && new != "" {
		return suppressSensitiveDiff
	}
	if suppressSensitiveEnvStr, ok := os.LookupEnv("AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF"); ok {
		if suppressDiff, err := strconv.ParseBool(suppressSensitiveEnvStr); err != nil {
			log.Printf("[ERROR] Invalid value for environment variable AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF. Accepted values 1, t, T, true, TRUE, True, 0, f, F, false, FALSE, False. %v", err.Error())
			suppressSensitiveDiff = false
		} else {
			suppressSensitiveDiff = suppressDiff
		}
	}
	return suppressSensitiveDiff
}
func validateInteger(val interface{}, key string) (warns []string, errs []error) {
	if _, err := strconv.ParseInt(val.(string), 10, 64); err != nil {
		errs = append(errs, fmt.Errorf("[ERROR] %q must be valid integers value: %v", key, val))
	}
	return
}
func validateBool(val interface{}, key string) (warns []string, errs []error) {
	if _, err := strconv.ParseBool(val.(string)); err != nil {
		errs = append(errs, fmt.Errorf("[ERROR] %q must be valid boolean value: %v", key, val))
	}
	return
}
func validateFloat(val interface{}, key string) (warns []string, errs []error) {
	if _, err := strconv.ParseFloat(val.(string), 64); err != nil {
		errs = append(errs, fmt.Errorf("[ERROR] %q must be valid float value: %v", key, val))
	}
	return
}
