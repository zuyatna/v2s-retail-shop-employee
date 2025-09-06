package domain

func CanListEmployees(accessLevel AccessLevel) bool {
	switch accessLevel {
	case AccessHR, AccessManager, AccessSupervisor:
		return true
	default:
		return false
	}
}
