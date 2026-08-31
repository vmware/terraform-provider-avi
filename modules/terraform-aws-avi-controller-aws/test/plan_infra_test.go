// +build unit

package test

import (
   "os"
   "github.com/gruntwork-io/terratest/modules/terraform"
   "github.com/stretchr/testify/assert"
   "testing"
)

func TestStaticSiteValidity(t *testing.T) {
   t.Parallel()
   clientID := os.Getenv("ARM_CLIENT_ID")
   clientSecret := os.Getenv("ARM_CLIENT_SECRET")
   subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
   tenantID := os.Getenv("ARM_TENANT_ID")

	if clientID, clientSecret, subscriptionID, tenantID == "" {
		t.Fatalf("Azure Credentials environment variables cannot be empty.")
	}
   _fixturesDir := "../examples/single-controller"
   terratestOptions := &terraform.Options{
      TerraformDir: _fixturesDir,
      Vars: map[string]interface{}{},
   }
   t.Logf("Running in %s", _fixturesDir)
   output := terraform.InitAndPlan(t, terratestOptions)
   assert.Contains(t, output, "Plan OK")
}