// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComputeBackendService_backendServiceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceBasicExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "backend-service%{random_suffix}"
  health_checks = ["${google_compute_http_health_check.default.self_link}"]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func testAccCheckComputeBackendServiceDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_backend_service" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/backendServices/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeBackendService still exists at %s", url)
		}
	}

	return nil
}
