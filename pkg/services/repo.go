// FIXME: golangci-lint
// nolint:revive
package services

import (
	"context"
	"errors"

	"github.com/redhatinsights/edge-api/pkg/db"
	"github.com/redhatinsights/edge-api/pkg/models"
	log "github.com/sirupsen/logrus"
)

// RepoServiceInterface defines the interface to handle the business logic of RHEL for Edge Devices
type RepoServiceInterface interface {
	GetRepoByID(repoID *uint) (*models.Repo, error)
}

// NewRepoService gives a instance of the main implementation of RepoServiceInterface
func NewRepoService(ctx context.Context, log *log.Entry) RepoServiceInterface {
	return &RepoService{
		ctx: ctx,
		log: log,
	}
}

// RepoService is the main implementation of a RepoServiceInterface
type RepoService struct {
	ctx context.Context
	log *log.Entry
}

// GetRepoByID receives RepoID uint and get a *models.Repo back
func (s *RepoService) GetRepoByID(repoID *uint) (*models.Repo, error) {
	if repoID == nil {
		s.log.Error("Image Repository is undefined")
		return nil, errors.New("image repository is undefined")
	}
	s.log.Debug("Retrieving repo by ID")
	var repo models.Repo
	result := db.DB.First(&repo, repoID)
	if result.Error != nil {
		s.log.Error("Error retrieving image repository")
		return nil, result.Error
	}
	s.log.Debug("Repo by ID retrieved successfully")
	return &repo, nil
}
