package datadog

func Int(val interface{}) *int {
	if ival, ok := val.(int); ok {
		return &ival
	}
	return nil
}
