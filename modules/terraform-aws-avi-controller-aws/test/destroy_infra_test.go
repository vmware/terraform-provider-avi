package test

import (
   "github.com/gruntwork-io/terratest/modules/terraform"
   test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
   "testing"
   "os"
   "strings"
)

func getTerraVars() map[string]interface{} {
   terraVars := make(map[string]interface{})

	for _, element := range os.Environ() {
      variable := strings.SplitN(element, "=", 2)

		if strings.HasPrefix(variable[0], "TF_VAR_") == true {
			envName := strings.Split(variable[0], "TF_VAR_")
			terraVars[envName[1]] = variable[1]
		}
	}
   return terraVars
}

func TestDeployment(t *testing.T) {
   t.Parallel()

   siteType := os.Getenv("site_type")

   if siteType == "" {
		t.Fatalf("site_type environment variable cannot be empty. single-site or gslb are valid values")
	}

   TerraformDir := "../examples/" + siteType

   // Uncomment these when doing local testing if you need to skip any stages.
   //os.Setenv("SKIP_destroy", "true")

   // Destroy the infrastructure
   test_structure.RunTestStage(t, "destroy", func() {
      terraformOptions := test_structure.LoadTerraformOptions(t, TerraformDir)
      terraform.Destroy(t, terraformOptions)
   })

}