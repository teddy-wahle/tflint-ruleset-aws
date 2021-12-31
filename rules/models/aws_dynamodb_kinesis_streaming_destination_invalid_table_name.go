// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule checks the pattern is valid
type AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule returns new rule with default attributes
func NewAwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule() *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule {
	return &AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule{
		resourceType:  "aws_dynamodb_kinesis_streaming_destination",
		attributeName: "table_name",
		max:           255,
		min:           3,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`),
	}
}

// Name returns the rule name
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule) Name() string {
	return "aws_dynamodb_kinesis_streaming_destination_invalid_table_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"table_name must be 255 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"table_name must be 3 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_.-]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}