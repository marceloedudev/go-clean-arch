package utils

// NewValidatorError validate
func NewValidatorError(err error) string {

	if err == nil {
		return "Internal server error"
	}

	return err.Error()

}
