package db

import (
	"context"

	"github.com/lenon/gofii/ent"
	"github.com/lenon/gofii/fnet"
)

type Connection struct {
	Client  *ent.Client
	Context context.Context
}

func Open(ctx context.Context, driver, dsn string) (*Connection, error) {
	client, err := ent.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	return &Connection{Context: ctx, Client: client}, nil
}

func (c *Connection) Close() error {
	return c.Client.Close()
}

func (c *Connection) Migrate() error {
	return c.Client.Schema.Create(c.Context)
}

func (c *Connection) UpsertDocument(doc *fnet.Document) error {
	referenceDate, err := doc.ParseReferenceDate()
	if err != nil {
		return err
	}

	submissionDate, err := doc.ParseSubmissionDate()
	if err != nil {
		return nil
	}

	categoryID, err := c.Client.FnetCategory.Create().
		SetName(doc.Category).
		OnConflict().
		UpdateNewValues().
		ID(c.Context)

	if err != nil {
		return err
	}

	create := c.Client.FnetDocument.Create().
		SetFnetID(doc.ID).
		SetAdditionalInformation(doc.AdditionalInformation).
		SetCategoryID(categoryID).
		SetCategoryStr(doc.Category).
		SetFundDescription(doc.FundDescription).
		SetFundMarketName(doc.FundMarketName).
		SetHighPriority(doc.HighPriority).
		SetReferenceDate(referenceDate).
		SetReferenceDateFormat(doc.ReferenceDateFormat).
		SetReferenceDateStr(doc.ReferenceDate).
		SetReviewed(doc.Reviewed).
		SetStatus(doc.Status).
		SetSubCategory1Str(doc.SubCategory1).
		SetSubCategory2Str(doc.SubCategory2).
		SetSubmissionDate(submissionDate).
		SetSubmissionDateStr(doc.SubmissionDate).
		SetSubmissionMethod(doc.SubmissionMethod).
		SetSubmissionMethodDescription(doc.SubmissionMethodDescription).
		SetSubmissionStatus(doc.SubmissionStatus).
		SetSubmissionStatusDescription(doc.SubmissionStatusDescription).
		SetVersion(doc.Version)

	if doc.SubCategory1 != "" {
		subCategory1ID, err := c.Client.FnetSubCategory1.Create().
			SetName(doc.SubCategory1).
			OnConflict().
			UpdateNewValues().
			ID(c.Context)

		if err != nil {
			return err
		}

		create.SetSubCategory1ID(subCategory1ID)
	}

	if doc.SubCategory2 != "" {
		subCategory2ID, err := c.Client.FnetSubCategory2.Create().
			SetName(doc.SubCategory2).
			OnConflict().
			UpdateNewValues().
			ID(c.Context)

		if err != nil {
			return err
		}

		create.SetSubCategory2ID(subCategory2ID)
	}

	return create.OnConflict().
		UpdateNewValues().
		Exec(c.Context)
}
