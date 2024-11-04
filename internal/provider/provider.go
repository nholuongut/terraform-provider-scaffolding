package provider

import (
	"github.com/nholuongutterraform-plugin-sdk/helper/schema"
	"github.com/nholuongutterraform-plugin-sdk/terraform"
)

func New() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"scaffolding_data_source": dataSourceScaffolding(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"scaffolding_resource": resourceScaffolding(),
		},
	}
}
