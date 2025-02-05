// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AvailabilityZone struct {
	Name *string `json:"name,omitempty"`
}

type BrokerEngineType struct {
	EngineType *string `json:"engineType,omitempty"`
}

type BrokerInstance struct {
	ConsoleURL *string   `json:"consoleURL,omitempty"`
	Endpoints  []*string `json:"endpoints,omitempty"`
	IPAddress  *string   `json:"ipAddress,omitempty"`
}

type BrokerInstanceOption struct {
	EngineType              *string   `json:"engineType,omitempty"`
	HostInstanceType        *string   `json:"hostInstanceType,omitempty"`
	StorageType             *string   `json:"storageType,omitempty"`
	SupportedEngineVersions []*string `json:"supportedEngineVersions,omitempty"`
}

type BrokerSummary struct {
	BrokerARN        *string      `json:"brokerARN,omitempty"`
	BrokerID         *string      `json:"brokerID,omitempty"`
	BrokerName       *string      `json:"brokerName,omitempty"`
	BrokerState      *string      `json:"brokerState,omitempty"`
	Created          *metav1.Time `json:"created,omitempty"`
	DeploymentMode   *string      `json:"deploymentMode,omitempty"`
	EngineType       *string      `json:"engineType,omitempty"`
	HostInstanceType *string      `json:"hostInstanceType,omitempty"`
}

type Configuration struct {
	ARN                    *string            `json:"arn,omitempty"`
	AuthenticationStrategy *string            `json:"authenticationStrategy,omitempty"`
	Created                *metav1.Time       `json:"created,omitempty"`
	Description            *string            `json:"description,omitempty"`
	EngineType             *string            `json:"engineType,omitempty"`
	EngineVersion          *string            `json:"engineVersion,omitempty"`
	ID                     *string            `json:"id,omitempty"`
	Name                   *string            `json:"name,omitempty"`
	Tags                   map[string]*string `json:"tags,omitempty"`
}

type ConfigurationID struct {
	ID       *string `json:"id,omitempty"`
	Revision *int64  `json:"revision,omitempty"`
}

type ConfigurationRevision struct {
	Created     *metav1.Time `json:"created,omitempty"`
	Description *string      `json:"description,omitempty"`
	Revision    *int64       `json:"revision,omitempty"`
}

type Configurations struct {
	Current *ConfigurationID   `json:"current,omitempty"`
	History []*ConfigurationID `json:"history,omitempty"`
	Pending *ConfigurationID   `json:"pending,omitempty"`
}

type EncryptionOptions struct {
	KMSKeyID       *string `json:"kmsKeyID,omitempty"`
	UseAWSOwnedKey *bool   `json:"useAWSOwnedKey,omitempty"`
}

type EngineVersion struct {
	Name *string `json:"name,omitempty"`
}

type LdapServerMetadataInput struct {
	Hosts                  []*string `json:"hosts,omitempty"`
	RoleBase               *string   `json:"roleBase,omitempty"`
	RoleName               *string   `json:"roleName,omitempty"`
	RoleSearchMatching     *string   `json:"roleSearchMatching,omitempty"`
	RoleSearchSubtree      *bool     `json:"roleSearchSubtree,omitempty"`
	ServiceAccountPassword *string   `json:"serviceAccountPassword,omitempty"`
	ServiceAccountUsername *string   `json:"serviceAccountUsername,omitempty"`
	UserBase               *string   `json:"userBase,omitempty"`
	UserRoleName           *string   `json:"userRoleName,omitempty"`
	UserSearchMatching     *string   `json:"userSearchMatching,omitempty"`
	UserSearchSubtree      *bool     `json:"userSearchSubtree,omitempty"`
}

type LdapServerMetadataOutput struct {
	Hosts                  []*string `json:"hosts,omitempty"`
	RoleBase               *string   `json:"roleBase,omitempty"`
	RoleName               *string   `json:"roleName,omitempty"`
	RoleSearchMatching     *string   `json:"roleSearchMatching,omitempty"`
	RoleSearchSubtree      *bool     `json:"roleSearchSubtree,omitempty"`
	ServiceAccountUsername *string   `json:"serviceAccountUsername,omitempty"`
	UserBase               *string   `json:"userBase,omitempty"`
	UserRoleName           *string   `json:"userRoleName,omitempty"`
	UserSearchMatching     *string   `json:"userSearchMatching,omitempty"`
	UserSearchSubtree      *bool     `json:"userSearchSubtree,omitempty"`
}

type Logs struct {
	Audit   *bool `json:"audit,omitempty"`
	General *bool `json:"general,omitempty"`
}

type LogsSummary struct {
	Audit           *bool        `json:"audit,omitempty"`
	AuditLogGroup   *string      `json:"auditLogGroup,omitempty"`
	General         *bool        `json:"general,omitempty"`
	GeneralLogGroup *string      `json:"generalLogGroup,omitempty"`
	Pending         *PendingLogs `json:"pending,omitempty"`
}

type PendingLogs struct {
	Audit   *bool `json:"audit,omitempty"`
	General *bool `json:"general,omitempty"`
}

type SanitizationWarning struct {
	AttributeName *string `json:"attributeName,omitempty"`
	ElementName   *string `json:"elementName,omitempty"`
}

type User struct {
	ConsoleAccess *bool     `json:"consoleAccess,omitempty"`
	Groups        []*string `json:"groups,omitempty"`
	Password      *string   `json:"password,omitempty"`
	Username      *string   `json:"username,omitempty"`
}

type UserPendingChanges struct {
	ConsoleAccess *bool     `json:"consoleAccess,omitempty"`
	Groups        []*string `json:"groups,omitempty"`
	PendingChange *string   `json:"pendingChange,omitempty"`
}

type UserSummary struct {
	PendingChange *string `json:"pendingChange,omitempty"`
	Username      *string `json:"username,omitempty"`
}

type WeeklyStartTime struct {
	DayOfWeek *string `json:"dayOfWeek,omitempty"`
	TimeOfDay *string `json:"timeOfDay,omitempty"`
	TimeZone  *string `json:"timeZone,omitempty"`
}
