package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/lenon/gofii/db"
	"github.com/lenon/gofii/fnet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func setupDB() (*db.Connection, error) {
	db, err := db.Open(context.Background(), "sqlite3", ":memory:?_fk=1")
	if err != nil {
		return nil, err
	}

	err = db.Migrate()
	if err != nil {
		return nil, err
	}

	return db, err
}

func TestUpsertNewDocument(t *testing.T) {
	db, err := setupDB()
	require.Nil(t, err)
	defer db.Close()

	document := &fnet.Document{
		ID:                          123,
		AdditionalInformation:       "my info",
		DocumentCategory:            "my category",
		DocumentStatus:              "AC",
		DocumentSubCategory1:        "my subcategory 1",
		DocumentSubCategory2:        "my subcategory 2",
		FundDescription:             "my fii",
		HighPriority:                false,
		MarketName:                  "my fii name",
		ReferenceDate:               "15/09/2022",
		ReferenceDateFormat:         "3",
		Reviewed:                    "N",
		Status:                      "AC",
		StatusDescription:           "Ativo com visualização",
		SubmissionDate:              "15/09/2022 20:30",
		SubmissionMethod:            "AP",
		SubmissionMethodDescription: "Apresentação",
		Version:                     2,
	}

	err = db.UpsertDocument(document)
	require.Nil(t, err)

	entity, err := db.Client.FnetDocument.Query().First(db.Context)
	require.Nil(t, err)

	assert.Equal(t, 1, entity.ID)
	assert.Equal(t, 123, entity.FnetID)
	assert.Equal(t, "my info", entity.AdditionalInformation)
	assert.Equal(t, "my category", entity.DocumentCategory)
	assert.Equal(t, "AC", entity.DocumentStatus)
	assert.Equal(t, "my subcategory 1", entity.DocumentSubCategory1)
	assert.Equal(t, "my subcategory 2", entity.DocumentSubCategory2)
	assert.Equal(t, "my fii", entity.FundDescription)
	assert.Equal(t, false, entity.HighPriority)
	assert.Equal(t, "my fii name", entity.MarketName)
	assert.Equal(t, "3", entity.ReferenceDateFormat)
	assert.True(t, time.Date(2022, 9, 15, 0, 0, 0, 0, time.Local).Equal(entity.ReferenceDate))
	assert.Equal(t, "15/09/2022", entity.ReferenceDateStr)
	assert.Equal(t, "N", entity.Reviewed)
	assert.Equal(t, "Ativo com visualização", entity.StatusDescription)
	assert.Equal(t, "AC", entity.Status)
	assert.True(t, time.Date(2022, 9, 15, 20, 30, 0, 0, time.Local).Equal(entity.SubmissionDate))
	assert.Equal(t, "15/09/2022 20:30", entity.SubmissionDateStr)
	assert.Equal(t, "Apresentação", entity.SubmissionMethodDescription)
	assert.Equal(t, "AP", entity.SubmissionMethod)
	assert.Equal(t, 2, entity.Version)
}

func TestUpsertAlreadyCreatedDocument(t *testing.T) {
	db, err := setupDB()
	require.Nil(t, err)
	defer db.Close()

	document := fnet.Document{
		ID:                          123,
		AdditionalInformation:       "my info",
		DocumentCategory:            "my category",
		DocumentStatus:              "AC",
		DocumentSubCategory1:        "my subcategory 1",
		DocumentSubCategory2:        "my subcategory 2",
		FundDescription:             "my fii",
		HighPriority:                false,
		MarketName:                  "my fii name",
		ReferenceDate:               "15/09/2022",
		ReferenceDateFormat:         "3",
		Reviewed:                    "N",
		Status:                      "AC",
		StatusDescription:           "Ativo com visualização",
		SubmissionDate:              "15/09/2022 20:30",
		SubmissionMethod:            "AP",
		SubmissionMethodDescription: "Apresentação",
		Version:                     2,
	}

	err = db.UpsertDocument(&document)
	require.Nil(t, err)

	documentCopy := document
	documentCopy.DocumentCategory = "my updated category"
	documentCopy.Version = 3

	// upsert the same fnet document again with updated fields
	err = db.UpsertDocument(&documentCopy)
	require.Nil(t, err)

	count := db.Client.FnetDocument.Query().CountX(db.Context)
	require.Equal(t, 1, count, "must upsert existing record and not create a new one")

	entity, err := db.Client.FnetDocument.Query().First(db.Context)
	require.Nil(t, err)

	assert.Equal(t, 1, entity.ID)
	assert.Equal(t, 123, entity.FnetID)
	assert.Equal(t, "my updated category", entity.DocumentCategory)
	assert.Equal(t, 3, entity.Version)
}

func TestUpsertInvalidDocument(t *testing.T) {
	db, err := setupDB()
	require.Nil(t, err)
	defer db.Close()

	document := fnet.Document{
		ID:                          123,
		AdditionalInformation:       "my info",
		DocumentCategory:            "",
		DocumentStatus:              "AC",
		DocumentSubCategory1:        "my subcategory 1",
		DocumentSubCategory2:        "my subcategory 2",
		FundDescription:             "",
		HighPriority:                false,
		MarketName:                  "my fii name",
		ReferenceDate:               "15/09/2022",
		ReferenceDateFormat:         "3",
		Reviewed:                    "N",
		Status:                      "AC",
		StatusDescription:           "Ativo com visualização",
		SubmissionDate:              "15/09/2022 20:30",
		SubmissionMethod:            "AP",
		SubmissionMethodDescription: "Apresentação",
		Version:                     2,
	}

	err = db.UpsertDocument(&document)
	assert.EqualError(t, err, "ent: validator failed for field \"FnetDocument.document_category\": value is less than the required length")
}