package resources

import "github.com/sl1pm4t/terraform-provider-kops/pkg/api/utils"

type RollingUpdateOptions struct {
	// Skip allows skipping cluster rolling update
	Skip bool
	utils.RollingUpdateOptions
}
