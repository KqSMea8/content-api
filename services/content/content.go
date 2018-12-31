package content

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zheng-ji/goSnowFlake"
	"os"
)

type ContentRepositoryInterface interface {
	// Put adds a new Greeting to the datastore
	FindOne(c context.Context, req FindOneRequest) *FindOneResponse
	Find(c context.Context, req FindRequest) *FindResponse
	Create(c context.Context, req CreateRequest) (*CreateResponse, error)
	Delete(c context.Context, req DeleteRequest) (*DeleteResponse, error)
	Update(c context.Context, req UpdateRequest) (*UpdateResponse, error)
}

type ContentRepository struct {
	db        *gorm.DB
	idBuilder *goSnowFlake.IdWorker
	ContentRepositoryInterface
}

func newContentRepository(db *gorm.DB) *ContentRepository {
	idBuilder, err := goSnowFlake.NewIdWorker(1)
	if err != nil {
		fmt.Printf("[services/content] Init snowFlake id_builder error: %+v", err)
		os.Exit(1)
	}
	return &ContentRepository{
		db:        db,
		idBuilder: idBuilder,
	}
}

// FindOneRequest ...
type FindOneRequest struct {
	ID        int64
	ContentID int64
}

// FindOneResponse ...
type FindOneResponse struct {
	DataInfo
}

// FindOne ...
func (cr ContentRepository) FindOne(c context.Context, req FindOneRequest) *FindOneResponse {
	var content PFContent
	cr.db.Where(&PFContent{
		ID:        req.ID,
		ContentID: req.ContentID,
	}).First(&content)
	res := &FindOneResponse{
		DataInfo: DataInfo{
			ID: content.ID,
		},
	}
	return res
}

// FindRequest ...
type FindRequest struct {
	ID        int64
	ContentID int64
}

// FindResponse ...
type FindResponse struct {
	List  []DataInfo
	Total int64
}

// FindOne ...
func (cr ContentRepository) Find(c context.Context, req FindOneRequest) *FindResponse {
	//var content PFContent
	cr.db.Where(&PFContent{
		ID:        req.ID,
		ContentID: req.ContentID,
	})
	res := &FindResponse{
		List:  []DataInfo{},
		Total: 0,
	}
	return res
}

// CreateRequest ...
type CreateRequest struct {
	Title       string
	Description string
	AuthorID    int64
	Category    string
	Type        int16
	Body        string
	Version     int16
	Extra       DataInfoExtra
}

// CreateResponse ...
type CreateResponse struct {
	DataInfo
}

// Create ...
func (cr ContentRepository) Create(c context.Context, req CreateRequest) (*CreateResponse, error) {
	// TODO: Duplicate title?

	// TODO: Param validating

	contentExtraStr, err := marshalContentExtraJson(&req.Extra)
	if err != nil {
		fmt.Printf("[services/content] Create: json marshal error: %+v", err)
		contentExtraStr, _ = marshalContentExtraJson(&DataInfoExtra{})
	}

	contentId, _ := cr.idBuilder.NextId()
	content := PFContent{
		ContentID:   contentId,
		Title:       req.Title,
		Description: req.Description,
		AuthorID:    req.AuthorID,
		Category:    req.Category,
		Type:        req.Type,
		Body:        req.Body,
		Version:     req.Version,
		Extra:       contentExtraStr,
	}
	if dbc := cr.db.Create(&content); dbc.Error != nil {
		fmt.Printf("[services/content] Create: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}

	responseExtra, err := UnmarshalContentExtraJson(content.Extra)
	if err != nil {
		fmt.Printf("[services/content] Create: UnmarshalContentJson error: %+v", err)
		responseExtra = &DataInfoExtra{}
	}
	res := &CreateResponse{
		DataInfo: DataInfo{
			ID:          content.ID,
			ContentID:   content.ContentID,
			Title:       content.Title,
			Description: content.Description,
			AuthorID:    content.AuthorID,
			Category:    content.Category,
			Type:        content.Type,
			Body:        content.Body,
			Version:     content.Version,
			CreatedAt:   content.CreatedAt,
			UpdatedAt:   content.UpdatedAt,
			Extra:       *responseExtra,
		},
	}
	return res, nil
}

// DeleteRequest ...
type DeleteRequest struct {
	ContentID int64
}

// DeleteResponse ...
type DeleteResponse struct {
	DataInfo
}

// Delete ...
func (cr ContentRepository) Delete(c context.Context, req DeleteRequest) (*DeleteResponse, error) {
	content := &PFContent{
		ContentID: req.ContentID,
	}
	if dbc := cr.db.Delete(content); dbc.Error != nil {
		fmt.Printf("[services/content] Delete: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	} else {
		fmt.Printf("%+v", dbc)
	}
	return &DeleteResponse{
		DataInfo: pfContentToData(*content),
	}, nil
}

// UpdateRequest ...
type UpdateRequest struct {
	Target PFContent

	Title       string
	Description string
	Category    string
	Type        int16
	Body        string
	//Extra       DataInfoExtra
}

// UpdateResponse ...
type UpdateResponse struct {
	DataInfo
}

// Update ...
func (cr ContentRepository) Update(c context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var target PFContent

	// TODO: 加上extra的update
	cr.db.Where("content_id = ?", req.Target.ContentID).Take(&target)
	var modified bool = false
	if req.Title != "" {
		target.Title = req.Title
		modified = true
	}
	if req.Description != "" {
		target.Description = req.Description
		modified = true
	}
	if req.Category != "" {
		target.Category = req.Category
		modified = true
	}
	if req.Type != 0 {
		target.Type = req.Type
		modified = true
	}
	if req.Body != "" {
		target.Body = req.Body
		modified = true
	}

	if modified {
		target.Version += 1
	} else {
		return nil, errors.New("NoContentModified")
	}
	if dbc := cr.db.Save(&target); dbc.Error != nil {
		fmt.Printf("[services/content] Update: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}

	return &UpdateResponse{
		DataInfo: pfContentToData(target),
	}, nil

}
