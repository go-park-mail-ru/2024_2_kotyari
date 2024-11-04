package utils

import "github.com/microcosm-cc/bluemonday"

type InputValidator struct {
	policy *bluemonday.Policy
}

func NewInputValidator() *InputValidator {
	return &InputValidator{
		policy: bluemonday.StrictPolicy(),
	}
}

func (v *InputValidator) SanitizeString(dirtyString string) string {
	return v.policy.Sanitize(dirtyString)
}
