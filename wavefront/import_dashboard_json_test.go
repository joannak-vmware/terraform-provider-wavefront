package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDashboardJson_Basic(t *testing.T) {
	resourceName := "wavefront_dashboard_json.json_foobar"
	var record wavefront.Dashboard

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontDashboardJSONDestroy,
		Steps: []resource.TestStep{
			{
				ExpectNonEmptyPlan: true,
				Config:             testAccCheckWavefrontDashboardJSONImporterBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontDashboardJSONExists("wavefront_dashboard_json.json_foobar", &record),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckWavefrontDashboardJSONImporterBasic() string {
	return `
resource "wavefront_dashboard_json" "json_foobar" {
	dashboard_json = <<EOF
{
  "name": "Terraform Test Dashboard Json",
  "description": "a",
  "eventFilterType": "BYCHART",
  "eventQuery": "",
  "defaultTimeWindow": "",
  "url": "tftestimport",
  "displayDescription": false,
  "displaySectionTableOfContents": true,
  "displayQueryParameters": false,
  "sections": [
    {
      "name": "section 1",
      "rows": [
        {
          "charts": [
            {
              "name": "chart 1",
              "sources": [
                {
                  "name": "source 1",
                  "query": "ts()",
                  "scatterPlotSource": "Y",
                  "querybuilderEnabled": false,
                  "sourceDescription": ""
                }
              ],
              "units": "someunit",
              "base": 0,
              "noDefaultEvents": false,
              "interpolatePoints": false,
              "includeObsoleteMetrics": false,
              "description": "This is chart 1, showing something",
              "chartSettings": {
                "type": "markdown-widget",
                "max": 100,
                "expectedDataSpacing": 120,
                "windowing": "full",
                "windowSize": 10,
                "autoColumnTags": false,
                "columnTags": "deprecated",
                "tagMode": "all",
                "numTags": 2,
                "customTags": [
                  "tag1",
                  "tag2"
                ],
                "groupBySource": true,
                "y1Max": 100,
                "y1Units": "units",
                "y0ScaleSIBy1024": true,
                "y1ScaleSIBy1024": true,
                "y0UnitAutoscaling": true,
                "y1UnitAutoscaling": true,
                "fixedLegendEnabled": true,
                "fixedLegendUseRawStats": true,
                "fixedLegendPosition": "RIGHT",
                "fixedLegendDisplayStats": [
                  "stat1",
                  "stat2"
                ],
                "fixedLegendFilterSort": "TOP",
                "fixedLegendFilterLimit": 1,
                "fixedLegendFilterField": "CURRENT",
                "plainMarkdownContent": "markdown content"
              },
              "chartAttributes": {
                "dashboardLayout": {
                  "x": 0,
                  "y": 0,
                  "w": 4,
                  "h": 7
                }
              },
              "summarization": "MEAN"
            }
          ],
          "heightFactor": 50
        }
      ]
    }
  ],
  "parameterDetails": {
    "param": {
      "hideFromView": false,
      "parameterType": "SIMPLE",
      "label": "test",
      "defaultValue": "Label",
      "valuesToReadableStrings": {
        "Label": "test"
      },
      "selectedLabel": "Label",
      "value": "test"
    }
  }
}
EOF
}
`
}
