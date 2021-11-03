// +build e2e

package test

import (
   "fmt"
   http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
   "github.com/gruntwork-io/terratest/modules/logger"
   "github.com/gruntwork-io/terratest/modules/retry"
   "github.com/gruntwork-io/terratest/modules/random"
   "github.com/gruntwork-io/terratest/modules/terraform"
   test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
   "github.com/stretchr/testify/assert"
   "strings"
   "testing"
   "time"
   "os"
   "crypto/tls"
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

func createRandomPrefix(base string) string {

   fmt.Println(base)
   random := strings.ToLower(random.UniqueId())
   prefix := fmt.Sprintf(base+"%s", random)
   fmt.Println(prefix)

   return prefix
}

func TestDeployment(t *testing.T) {
   t.Parallel()

   siteType := os.Getenv("site_type")

   if siteType == "" {
		t.Fatalf("site_type environment variable cannot be empty. single-site or gslb are valid values")
	}
   
   TerraformDir := "../examples/" + siteType

   // Uncomment these when doing local testing if you need to skip any stages.
   //os.Setenv("SKIP_bootstrap", "true")
   //os.Setenv("SKIP_apply", "true")
   os.Setenv("SKIP_perpetual_diff", "true")
   //os.Setenv("SKIP_singlesite_controller_tests", "true")
   //os.Setenv("SKIP_gslb_controller_tests", "true")
   os.Setenv("SKIP_teardown", "true")
   os.Setenv("SKIP_destroy", "true")

   switch siteType {
      case "single-site":
         os.Setenv("SKIP_gslb_controller_tests", "true")
      case "gslb":
         os.Setenv("SKIP_singlesite_controller_tests", "true")
   }

   test_structure.RunTestStage(t, "bootstrap", func() {

      switch siteType {
         case "single-site":
            stringName := "TF_VAR_name_prefix"
            randomName := createRandomPrefix("terraform")
            os.Setenv(stringName, randomName)
            test_structure.SaveString(t, TerraformDir, stringName, randomName )
         case "gslb":
            stringNameEast := "TF_VAR_name_prefix_east"
            stringNameWest := "TF_VAR_name_prefix_west"
            prefixEast := createRandomPrefix("tfeast")
            prefixWest := createRandomPrefix("tfwest")
            os.Setenv(stringNameEast, prefixEast)
            os.Setenv(stringNameWest, prefixWest)
            test_structure.SaveString(t, TerraformDir, stringNameEast, prefixEast )
            test_structure.SaveString(t, TerraformDir, stringNameWest, prefixWest )
         }
   })

   // At the end of the test, run `terraform destroy` to clean up any resources that were created
   defer test_structure.RunTestStage(t, "teardown", func() {
      terraformOptions := test_structure.LoadTerraformOptions(t, TerraformDir)
      terraform.Destroy(t, terraformOptions)
   })

   // Apply the infrastructure
   test_structure.RunTestStage(t, "apply", func() {
      terraVars := getTerraVars()
      terratestOptions := &terraform.Options{
         // The path to where your Terraform code is located
         TerraformDir: TerraformDir,
         Vars: terraVars,
      }
      // Save the terraform oprions for future reference
      test_structure.SaveTerraformOptions(t, TerraformDir, terratestOptions)
      output := terraform.InitAndApply(t, terratestOptions)
      assert.Contains(t, output, "Apply complete!")
   })

   // Run perpetual diff
   test_structure.RunTestStage(t, "perpetual_diff", func() {
      terraformOptions := test_structure.LoadTerraformOptions(t, TerraformDir)
      planResult := terraform.Plan(t, terraformOptions)

      // Make sure the plan shows zero changes
      assert.Contains(t, planResult, "No changes.")
   })

   // Destroy the infrastructure - used during development - the teardown function will be used normally
   test_structure.RunTestStage(t, "destroy", func() {
      var targets []string
      targets = append(targets, "azurerm_resource_group.avi")
      terraVars := getTerraVars()
      terraformOptions := &terraform.Options{
         // The path to where your Terraform code is located
         TerraformDir: TerraformDir,
         Vars: terraVars,
         Targets: targets,
      }
      terraform.Destroy(t, terraformOptions)
   })

   // Run Controller tests for a single site deployment
   test_structure.RunTestStage(t, "singlesite_controller_tests", func() {
      var controllerIPs []string
      terraformOptions := test_structure.LoadTerraformOptions(t, TerraformDir)
      controllerInfo :=  terraform.OutputListOfObjects(t, terraformOptions, "controllers" )
      
      publicIP := os.Getenv("TF_VAR_controller_public_address")

      switch publicIP {
         case "true":
            for index, value := range controllerInfo {
               _ = index
               IP := value["public_ip_address"]
               publicIP := fmt.Sprintf("%v", IP)
               controllerIPs = append(controllerIPs, publicIP)
            }
         case "false":
            for index, value := range controllerInfo {
               _ = index
               IP := value["private_ip_address"]
               privateIP := fmt.Sprintf("%v", IP)
               controllerIPs = append(controllerIPs, privateIP)
            }
      }
      for index, controller := range controllerIPs {
         _ = index
         testURL(t, controller, "", 200, "Avi Vantage Controller")
         testURL(t, controller, "notfound", 404, "not found")
      }
   })

   // Run Controller tests for a GSLB deployment
   test_structure.RunTestStage(t, "gslb_controller_tests", func() {
      var controllerIPs []string
      terraformOptions := test_structure.LoadTerraformOptions(t, TerraformDir)

      controllersEast :=  terraform.OutputListOfObjects(t, terraformOptions, "controllers_east" )
      controllersWest :=  terraform.OutputListOfObjects(t, terraformOptions, "controllers_west" )

      publicIP := os.Getenv("TF_VAR_controller_public_address")

      switch publicIP {
         case "true":
            for index, value := range controllersEast {
               _ = index
               IP := value["public_ip_address"]
               publicIP := fmt.Sprintf("%v", IP)
               controllerIPs = append(controllerIPs, publicIP)
            }
            for index, value := range controllersWest {
               _ = index
               IP := value["public_ip_address"]
               publicIP := fmt.Sprintf("%v", IP)
               controllerIPs = append(controllerIPs, publicIP)
            }
         case "false":
            for index, value := range controllersEast {
               _ = index
               IP := value["private_ip_address"]
               privateIP := fmt.Sprintf("%v", IP)
               controllerIPs = append(controllerIPs, privateIP)
            }
            for index, value := range controllersWest {
               _ = index
               IP := value["private_ip_address"]
               privateIP := fmt.Sprintf("%v", IP)
               controllerIPs = append(controllerIPs, privateIP)
            }
      }
      for index, controller := range controllerIPs {
         _ = index
         testURL(t, controller, "", 200, "Avi Vantage Controller")
         testURL(t, controller, "notfound", 404, "not found")
      }
      
   })

}

func testURL(t *testing.T, endpoint string, path string, expectedStatus int, expectedBody string) {
   tlsConfig := tls.Config{InsecureSkipVerify: true}
   url := fmt.Sprintf("%s://%s/%s", "https", endpoint, path)
   actionDescription := fmt.Sprintf("Calling %s", url)
   output := retry.DoWithRetry(t, actionDescription, 24 , 30 * time.Second, func() (string, error) {
      statusCode, body := http_helper.HttpGet(t, url, &tlsConfig)
      if statusCode == expectedStatus {
         logger.Logf(t, "Got expected status code %d from URL %s", expectedStatus, url)
         return body, nil
      }
      return "", fmt.Errorf("Recieved status %d from %s. Retrying...", statusCode, url)
   })
   assert.Contains(t, output, expectedBody, "Body should contain expected text")
}