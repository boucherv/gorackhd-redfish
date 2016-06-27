package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*AccountService100AccountService This is the schema definition for the Account Service.  It represents the properties for the service itself and has links to the actual list of accounts.

swagger:model AccountService.1.0.0_AccountService
*/
type AccountService100AccountService struct {

	/* at odata context
	 */
	AtOdataContext Odata400Context `json:"@odata.context,omitempty"`

	/* at odata id
	 */
	AtOdataID Odata400ID `json:"@odata.id,omitempty"`

	/* at odata type
	 */
	AtOdataType Odata400Type `json:"@odata.type,omitempty"`

	/* The interval of time since the last failed login attempt at which point the lockout threshold counter for the account is reset to zero. Must be less than or equal to AccountLockoutDuration

	Minimum: 0
	*/
	AccountLockoutCounterResetAfter *float64 `json:"AccountLockoutCounterResetAfter,omitempty"`

	/* The time an account is locked after the account lockout threshold is met. Must be >= AccountLockoutResetAfter. If set to 0, no lockout will occur.

	Minimum: 0
	*/
	AccountLockoutDuration *float64 `json:"AccountLockoutDuration,omitempty"`

	/* The number of failed login attempts before a user account is locked for a specified duration. (0=never locked)

	Minimum: 0
	*/
	AccountLockoutThreshold *float64 `json:"AccountLockoutThreshold,omitempty"`

	/* Link to a collection of Manager Accounts

	Read Only: true
	*/
	Accounts *ManagerAccountCollectionManagerAccountCollection `json:"Accounts,omitempty"`

	/* This is the number of authorization failures that need to occur before the failure attempt is logged to the manager log.

	Minimum: 0
	*/
	AuthFailureLoggingThreshold *float64 `json:"AuthFailureLoggingThreshold,omitempty"`

	/* description
	 */
	Description ResourceDescription `json:"Description,omitempty"`

	/* Id
	 */
	ID ResourceID `json:"Id,omitempty"`

	/* This is the maximum password length for this service.

	Read Only: true
	Minimum: 0
	*/
	MaxPasswordLength float64 `json:"MaxPasswordLength,omitempty"`

	/* This is the minimum password length for this service.

	Read Only: true
	Minimum: 0
	*/
	MinPasswordLength float64 `json:"MinPasswordLength,omitempty"`

	/* name
	 */
	Name ResourceName `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* Link to a collection of Roles

	Read Only: true
	*/
	Roles *RoleCollectionRoleCollection `json:"Roles,omitempty"`

	/* This indicates whether this service is enabled.
	 */
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	/* status
	 */
	Status *ResourceStatus `json:"Status,omitempty"`
}

// Validate validates this account service 1 0 0 account service
func (m *AccountService100AccountService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountLockoutCounterResetAfter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAccountLockoutDuration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAccountLockoutThreshold(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAuthFailureLoggingThreshold(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaxPasswordLength(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinPasswordLength(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountService100AccountService) validateAccountLockoutCounterResetAfter(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountLockoutCounterResetAfter) { // not required
		return nil
	}

	if err := validate.Minimum("AccountLockoutCounterResetAfter", "body", float64(*m.AccountLockoutCounterResetAfter), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountService100AccountService) validateAccountLockoutDuration(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountLockoutDuration) { // not required
		return nil
	}

	if err := validate.Minimum("AccountLockoutDuration", "body", float64(*m.AccountLockoutDuration), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountService100AccountService) validateAccountLockoutThreshold(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountLockoutThreshold) { // not required
		return nil
	}

	if err := validate.Minimum("AccountLockoutThreshold", "body", float64(*m.AccountLockoutThreshold), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountService100AccountService) validateAuthFailureLoggingThreshold(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthFailureLoggingThreshold) { // not required
		return nil
	}

	if err := validate.Minimum("AuthFailureLoggingThreshold", "body", float64(*m.AuthFailureLoggingThreshold), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountService100AccountService) validateMaxPasswordLength(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxPasswordLength) { // not required
		return nil
	}

	if err := validate.Minimum("MaxPasswordLength", "body", float64(m.MaxPasswordLength), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountService100AccountService) validateMinPasswordLength(formats strfmt.Registry) error {

	if swag.IsZero(m.MinPasswordLength) { // not required
		return nil
	}

	if err := validate.Minimum("MinPasswordLength", "body", float64(m.MinPasswordLength), 0, false); err != nil {
		return err
	}

	return nil
}
