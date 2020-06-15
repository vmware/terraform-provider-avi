package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func suppressPassphraseDiffs(k, old, new string, d *schema.ResourceData) bool {
	if d.Get("changed").(bool) {
		return false
	}
	return true
}

func suppressChangedDiffs(k, old, new string, d *schema.ResourceData) bool {
	return old == "true" && new == "false"
}
