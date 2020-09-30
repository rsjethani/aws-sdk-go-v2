// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// The resource being created already exists.
type AlreadyExistsException struct {
	Message *string
}

func (e *AlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsException) ErrorCode() string             { return "AlreadyExistsException" }
func (e *AlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *AlreadyExistsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *AlreadyExistsException) HasMessage() bool {
	return e.Message != nil
}

// There is a concurrent modification of resources.
type ConcurrentModificationException struct {
	Message *string
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ConcurrentModificationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ConcurrentModificationException) HasMessage() bool {
	return e.Message != nil
}

// The request failed because this device is no longer registered and therefore no
// longer managed by this account.
type DeviceNotRegisteredException struct {
	Message *string
}

func (e *DeviceNotRegisteredException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeviceNotRegisteredException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeviceNotRegisteredException) ErrorCode() string             { return "DeviceNotRegisteredException" }
func (e *DeviceNotRegisteredException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DeviceNotRegisteredException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DeviceNotRegisteredException) HasMessage() bool {
	return e.Message != nil
}

// The Certificate Authority can't issue or revoke a certificate.
type InvalidCertificateAuthorityException struct {
	Message *string
}

func (e *InvalidCertificateAuthorityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCertificateAuthorityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCertificateAuthorityException) ErrorCode() string {
	return "InvalidCertificateAuthorityException"
}
func (e *InvalidCertificateAuthorityException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *InvalidCertificateAuthorityException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidCertificateAuthorityException) HasMessage() bool {
	return e.Message != nil
}

// The device is in an invalid state.
type InvalidDeviceException struct {
	Message *string
}

func (e *InvalidDeviceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDeviceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDeviceException) ErrorCode() string             { return "InvalidDeviceException" }
func (e *InvalidDeviceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidDeviceException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidDeviceException) HasMessage() bool {
	return e.Message != nil
}

// A password in SecretsManager is in an invalid state.
type InvalidSecretsManagerResourceException struct {
	Message *string
}

func (e *InvalidSecretsManagerResourceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSecretsManagerResourceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSecretsManagerResourceException) ErrorCode() string {
	return "InvalidSecretsManagerResourceException"
}
func (e *InvalidSecretsManagerResourceException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *InvalidSecretsManagerResourceException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidSecretsManagerResourceException) HasMessage() bool {
	return e.Message != nil
}

// The service linked role is locked for deletion.
type InvalidServiceLinkedRoleStateException struct {
	Message *string
}

func (e *InvalidServiceLinkedRoleStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidServiceLinkedRoleStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidServiceLinkedRoleStateException) ErrorCode() string {
	return "InvalidServiceLinkedRoleStateException"
}
func (e *InvalidServiceLinkedRoleStateException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *InvalidServiceLinkedRoleStateException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidServiceLinkedRoleStateException) HasMessage() bool {
	return e.Message != nil
}

// The attempt to update a user is invalid due to the user's current status.
type InvalidUserStatusException struct {
	Message *string
}

func (e *InvalidUserStatusException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidUserStatusException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidUserStatusException) ErrorCode() string             { return "InvalidUserStatusException" }
func (e *InvalidUserStatusException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidUserStatusException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidUserStatusException) HasMessage() bool {
	return e.Message != nil
}

// You are performing an action that would put you beyond your account's limits.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// The name sent in the request is already in use.
type NameInUseException struct {
	Message *string
}

func (e *NameInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NameInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NameInUseException) ErrorCode() string             { return "NameInUseException" }
func (e *NameInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *NameInUseException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NameInUseException) HasMessage() bool {
	return e.Message != nil
}

// The resource is not found.
type NotFoundException struct {
	Message *string
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *NotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NotFoundException) HasMessage() bool {
	return e.Message != nil
}

// Another resource is associated with the resource in the request.
type ResourceAssociatedException struct {
	Message *string
}

func (e *ResourceAssociatedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAssociatedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAssociatedException) ErrorCode() string             { return "ResourceAssociatedException" }
func (e *ResourceAssociatedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceAssociatedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceAssociatedException) HasMessage() bool {
	return e.Message != nil
}

// The resource in the request is already in use.
type ResourceInUseException struct {
	Message *string

	ClientRequestToken *string
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string             { return "ResourceInUseException" }
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceInUseException) GetClientRequestToken() string {
	return ptr.ToString(e.ClientRequestToken)
}
func (e *ResourceInUseException) HasClientRequestToken() bool {
	return e.ClientRequestToken != nil
}
func (e *ResourceInUseException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceInUseException) HasMessage() bool {
	return e.Message != nil
}

// The skill must be linked to a third-party account.
type SkillNotLinkedException struct {
	Message *string
}

func (e *SkillNotLinkedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SkillNotLinkedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SkillNotLinkedException) ErrorCode() string             { return "SkillNotLinkedException" }
func (e *SkillNotLinkedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *SkillNotLinkedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *SkillNotLinkedException) HasMessage() bool {
	return e.Message != nil
}

// The caller has no permissions to operate on the resource involved in the API
// call.
type UnauthorizedException struct {
	Message *string
}

func (e *UnauthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedException) ErrorCode() string             { return "UnauthorizedException" }
func (e *UnauthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *UnauthorizedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *UnauthorizedException) HasMessage() bool {
	return e.Message != nil
}