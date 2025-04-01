package handler

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	var errors []string

	if r.Role == "" {
		errors = append(errors, errParamIsRequired("role", "string").Error())
	}
	if r.Company == "" {
		errors = append(errors, errParamIsRequired("company", "string").Error())
	}
	if r.Location == "" {
		errors = append(errors, errParamIsRequired("location", "string").Error())
	}
	if r.Link == "" {
		errors = append(errors, errParamIsRequired("link", "string").Error())
	}
	if r.Salary <= 0 {
		errors = append(errors, errParamIsRequired("salary", "int64").Error())
	}
	if r.Remote == nil {
		errors = append(errors, errParamIsRequired("remote", "bool").Error())
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided validation is truthy
	if r.Role != "" ||
		r.Company != "" ||
		r.Location != "" ||
		r.Remote != nil ||
		r.Link != "" ||
		r.Salary > 0 {
		return nil
	}

	// If none of the fields were provided, return fals
	return fmt.Errorf("at least one valid field must be provided")
}
