// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodestarconnectionsConnectionInvalidHostArnRule checks the pattern is valid
type AwsCodestarconnectionsConnectionInvalidHostArnRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsCodestarconnectionsConnectionInvalidHostArnRule returns new rule with default attributes
func NewAwsCodestarconnectionsConnectionInvalidHostArnRule() *AwsCodestarconnectionsConnectionInvalidHostArnRule {
	return &AwsCodestarconnectionsConnectionInvalidHostArnRule{
		resourceType:  "aws_codestarconnections_connection",
		attributeName: "host_arn",
		max:           256,
		pattern:       regexp.MustCompile(`^arn:aws(-[\w]+)*:codestar-connections:.+:[0-9]{12}:host\/.+$`),
	}
}

// Name returns the rule name
func (r *AwsCodestarconnectionsConnectionInvalidHostArnRule) Name() string {
	return "aws_codestarconnections_connection_invalid_host_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodestarconnectionsConnectionInvalidHostArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodestarconnectionsConnectionInvalidHostArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodestarconnectionsConnectionInvalidHostArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodestarconnectionsConnectionInvalidHostArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"host_arn must be 256 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws(-[\w]+)*:codestar-connections:.+:[0-9]{12}:host\/.+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
