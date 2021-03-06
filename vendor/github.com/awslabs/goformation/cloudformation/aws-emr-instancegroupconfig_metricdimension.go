package cloudformation

import (
	"encoding/json"
)

// AWSEMRInstanceGroupConfig_MetricDimension AWS CloudFormation Resource (AWS::EMR::InstanceGroupConfig.MetricDimension)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html
type AWSEMRInstanceGroupConfig_MetricDimension struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html#cfn-elasticmapreduce-instancegroupconfig-metricdimension-key
	Key *Value `json:"Key,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html#cfn-elasticmapreduce-instancegroupconfig-metricdimension-value
	Value *Value `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_MetricDimension) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.MetricDimension"
}

func (r *AWSEMRInstanceGroupConfig_MetricDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
