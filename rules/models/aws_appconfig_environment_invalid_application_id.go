// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppconfigEnvironmentInvalidApplicationIDRule checks the pattern is valid
type AwsAppconfigEnvironmentInvalidApplicationIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsAppconfigEnvironmentInvalidApplicationIDRule returns new rule with default attributes
func NewAwsAppconfigEnvironmentInvalidApplicationIDRule() *AwsAppconfigEnvironmentInvalidApplicationIDRule {
	return &AwsAppconfigEnvironmentInvalidApplicationIDRule{
		resourceType:  "aws_appconfig_environment",
		attributeName: "application_id",
		pattern:       regexp.MustCompile(`^[a-z0-9]{4,7}$`),
	}
}

// Name returns the rule name
func (r *AwsAppconfigEnvironmentInvalidApplicationIDRule) Name() string {
	return "aws_appconfig_environment_invalid_application_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppconfigEnvironmentInvalidApplicationIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppconfigEnvironmentInvalidApplicationIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppconfigEnvironmentInvalidApplicationIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppconfigEnvironmentInvalidApplicationIDRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9]{4,7}$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
