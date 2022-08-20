package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/models"
	"gorm.io/gorm"
)

type Domain interface {
	// Create adds a new domain to the database
	Create(ctx *gin.Context, domain models.Domain) (models.Domain, error)
	// Read gets a domain from the database
	Read(ctx *gin.Context, domain models.Domain) (models.Domain, error)
	// ReadAll gets all domains from the database
	ReadAll(ctx *gin.Context) ([]models.Domain, error)
	// Update updates a domain in the database
	Update(ctx *gin.Context, domain *models.Domain) error
	// Delete removes a domain from the database
	Delete(ctx *gin.Context, domain *models.Domain) error
}

type domain struct {
	db gorm.DB
}

// Create adds a new domain to the database
func (repo *domain) Create(ctx *gin.Context, domain models.Domain) (models.Domain, error) {
	return domain, nil
}

// Read gets a domain from the database
func (repo *domain) Read(ctx *gin.Context, domain models.Domain) (models.Domain, error) {
	return domain, nil
}

// ReadAll gets all domains from the database
func (repo *domain) ReadAll(ctx *gin.Context) ([]models.Domain, error) {
	return nil, nil
}

// Update updates a domain in the database
func (repo *domain) Update(ctx *gin.Context, domain *models.Domain) error {
	return nil
}

// Delete removes a domain from the database
func (repo *domain) Delete(ctx *gin.Context, domain *models.Domain) error {
	return nil
}

// NewDomain returns a new domain repository
func NewDomain(db *gorm.DB) Domain {
	return &domain{db: *db}
}
