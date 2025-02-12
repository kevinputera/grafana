package api

import (
	"time"

	"github.com/grafana/grafana/pkg/services/cloudmigration"
)

// swagger:parameters getCloudMigrationToken
type GetCloudMigrationToken struct {
}

// swagger:response cloudMigrationGetTokenResponse
type CloudMigrationGetTokenResponse struct {
	// in: body
	Body GetAccessTokenResponseDTO
}

type GetAccessTokenResponseDTO struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	ExpiresAt   string `json:"expiresAt"`
	FirstUsedAt string `json:"firstUsedAt"`
	LastUsedAt  string `json:"lastUsedAt"`
	CreatedAt   string `json:"createdAt"`
}

// swagger:response cloudMigrationCreateTokenResponse
type CloudMigrationCreateTokenResponse struct {
	// in: body
	Body CreateAccessTokenResponseDTO
}

type CreateAccessTokenResponseDTO struct {
	Token string `json:"token"`
}

// swagger:parameters deleteCloudMigrationToken
type DeleteCloudMigrationToken struct {
	// UID of a cloud migration token
	// in: path
	UID string `json:"uid"`
}

// swagger:response cloudMigrationDeleteTokenResponse
type CloudMigrationDeleteTokenResponse struct {
}

// swagger:response cloudMigrationSessionListResponse
type CloudMigrationSessionListResponse struct {
	// in: body
	Body CloudMigrationSessionListResponseDTO
}

type CloudMigrationSessionResponseDTO struct {
	UID     string    `json:"uid"`
	Slug    string    `json:"slug"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type CloudMigrationSessionListResponseDTO struct {
	Sessions []CloudMigrationSessionResponseDTO `json:"sessions"`
}

// swagger:parameters getSession
type GetCloudMigrationSessionRequest struct {
	// UID of a migration session
	// in: path
	UID string `json:"uid"`
}

// swagger:response cloudMigrationSessionResponse
type CloudMigrationSessionResponse struct {
	// in: body
	Body CloudMigrationSessionResponseDTO
}

// swagger:parameters createSession
type CreateSession struct {
	// in:body
	// required:true
	Body CloudMigrationSessionRequestDTO
}

type CloudMigrationSessionRequestDTO struct {
	AuthToken string `json:"authToken"`
}

// swagger:parameters runCloudMigration
type RunCloudMigrationRequest struct {
	// UID of a migration
	// in: path
	UID string `json:"uid"`
}

// swagger:response cloudMigrationRunResponse
type CloudMigrationRunResponse struct {
	// in: body
	Body MigrateDataResponseDTO
}

type MigrateDataResponseDTO struct {
	RunUID string                       `json:"uid"`
	Items  []MigrateDataResponseItemDTO `json:"items"`
}

type MigrateDataResponseItemDTO struct {
	// required:true
	Type MigrateDataType `json:"type"`
	// required:true
	RefID string `json:"refId"`
	// required:true
	Status ItemStatus `json:"status"`
	Error  string     `json:"error,omitempty"`
}

// swagger:enum MigrateDataType
type MigrateDataType string

const (
	DashboardDataType  MigrateDataType = "DASHBOARD"
	DatasourceDataType MigrateDataType = "DATASOURCE"
	FolderDataType     MigrateDataType = "FOLDER"
)

// swagger:enum ItemStatus
type ItemStatus string

const (
	ItemStatusOK    ItemStatus = "OK"
	ItemStatusError ItemStatus = "ERROR"
)

// swagger:parameters getCloudMigrationRun
type GetMigrationRunParams struct {
	// RunUID of a migration run
	// in: path
	RunUID string `json:"runUID"`
}

// swagger:parameters getCloudMigrationRunList
type GetCloudMigrationRunList struct {
	// UID of a migration
	// in: path
	UID string `json:"uid"`
}

// swagger:response cloudMigrationRunListResponse
type CloudMigrationRunListResponse struct {
	// in: body
	Body CloudMigrationRunListDTO
}

type CloudMigrationRunListDTO struct {
	Runs []MigrateDataResponseListDTO `json:"runs"`
}

type MigrateDataResponseListDTO struct {
	RunUID string `json:"uid"`
}

// swagger:parameters deleteSession
type DeleteMigrationSessionRequest struct {
	// UID of a migration session
	// in: path
	UID string `json:"uid"`
}

// utility funcs for converting to/from DTO

func convertSessionListToDTO(sl cloudmigration.CloudMigrationSessionListResponse) CloudMigrationSessionListResponseDTO {
	slDTOs := make([]CloudMigrationSessionResponseDTO, len(sl.Sessions))
	for i := 0; i < len(slDTOs); i++ {
		s := sl.Sessions[i]
		slDTOs[i] = CloudMigrationSessionResponseDTO{
			UID:     s.UID,
			Slug:    s.Slug,
			Created: s.Created,
			Updated: s.Updated,
		}
	}
	return CloudMigrationSessionListResponseDTO{
		Sessions: slDTOs,
	}
}

func convertMigrateDataResponseToDTO(r cloudmigration.MigrateDataResponse) MigrateDataResponseDTO {
	items := make([]MigrateDataResponseItemDTO, len(r.Items))
	for i := 0; i < len(r.Items); i++ {
		item := r.Items[i]
		items[i] = MigrateDataResponseItemDTO{
			Type:   MigrateDataType(item.Type),
			RefID:  item.RefID,
			Status: ItemStatus(item.Status),
			Error:  item.Error,
		}
	}
	return MigrateDataResponseDTO{
		RunUID: r.RunUID,
		Items:  items,
	}
}

// Base snapshot without results
type SnapshotDTO struct {
	SnapshotUID string         `json:"uid"`
	Status      SnapshotStatus `json:"status"`
	SessionUID  string         `json:"sessionUid"`
	Created     time.Time      `json:"created"`
	Finished    time.Time      `json:"finished"`
}

// swagger:enum SnapshotStatus
type SnapshotStatus string

const (
	SnapshotStatusInitializing      SnapshotStatus = "INITIALIZING"
	SnapshotStatusCreating          SnapshotStatus = "CREATING"
	SnapshotStatusPendingUpload     SnapshotStatus = "PENDING_UPLOAD"
	SnapshotStatusUploading         SnapshotStatus = "UPLOADING"
	SnapshotStatusPendingProcessing SnapshotStatus = "PENDING_PROCESSING"
	SnapshotStatusProcessing        SnapshotStatus = "PROCESSING"
	SnapshotStatusFinished          SnapshotStatus = "FINISHED"
	SnapshotStatusError             SnapshotStatus = "ERROR"
	SnapshotStatusUnknown           SnapshotStatus = "UNKNOWN"
)

func fromSnapshotStatus(status cloudmigration.SnapshotStatus) SnapshotStatus {
	switch status {
	case cloudmigration.SnapshotStatusInitializing:
		return SnapshotStatusInitializing
	case cloudmigration.SnapshotStatusCreating:
		return SnapshotStatusCreating
	case cloudmigration.SnapshotStatusPendingUpload:
		return SnapshotStatusPendingUpload
	case cloudmigration.SnapshotStatusUploading:
		return SnapshotStatusUploading
	case cloudmigration.SnapshotStatusPendingProcessing:
		return SnapshotStatusPendingProcessing
	case cloudmigration.SnapshotStatusProcessing:
		return SnapshotStatusProcessing
	case cloudmigration.SnapshotStatusFinished:
		return SnapshotStatusFinished
	case cloudmigration.SnapshotStatusError:
		return SnapshotStatusError
	default:
		return SnapshotStatusUnknown
	}
}

// swagger:parameters createSnapshot
type CreateSnapshotRequest struct {
	// UID of a session
	// in: path
	UID string `json:"uid"`
}

// swagger:response createSnapshotResponse
type CreateSnapshotResponse struct {
	// in: body
	Body CreateSnapshotResponseDTO
}

type CreateSnapshotResponseDTO struct {
	SnapshotUID string `json:"uid"`
}

// swagger:parameters getSnapshot
type GetSnapshotParams struct {
	// Session UID of a session
	// in: path
	UID string `json:"uid"`

	// UID of a snapshot
	// in: path
	SnapshotUID string `json:"snapshotUid"`
}

// swagger:response getSnapshotResponse
type GetSnapshotResponse struct {
	// in: body
	Body GetSnapshotResponseDTO
}

type GetSnapshotResponseDTO struct {
	SnapshotDTO
	Results []MigrateDataResponseItemDTO `json:"results"`
}

// swagger:parameters getShapshotList
type GetSnapshotListParams struct {
	// Offset is used for pagination with limit
	// in:query
	// required:false
	// default: 0
	Offset int `json:"offset"`
	// Max limit for results returned.
	// in:query
	// required:false
	// default: 100
	Limit int `json:"limit"`
	// Session UID of a session
	// in: path
	UID string `json:"uid"`
}

// swagger:response snapshotListResponse
type SnapshotListResponse struct {
	// in: body
	Body SnapshotListResponseDTO
}

type SnapshotListResponseDTO struct {
	Snapshots []SnapshotDTO `json:"snapshots"`
}

// swagger:parameters uploadSnapshot
type UploadSnapshotParams struct {
	// Session UID of a session
	// in: path
	UID string `json:"uid"`

	// UID of a snapshot
	// in: path
	SnapshotUID string `json:"snapshotUid"`
}

// swagger:parameters cancelSnapshot
type CancelSnapshotParams struct {
	// Session UID of a session
	// in: path
	UID string `json:"uid"`

	// UID of a snapshot
	// in: path
	SnapshotUID string `json:"snapshotUid"`
}
