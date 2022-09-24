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
	referenceDate, err := doc.ParsedReferenceDate()
	if err != nil {
		return err
	}

	submissionDate, err := doc.ParsedSubmissionDate()
	if err != nil {
		return nil
	}

	category_id, err := c.Client.FnetCategory.Create().
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
		SetCategoryID(category_id).
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
		sub_category1_id, err := c.Client.FnetSubCategory1.Create().
			SetName(doc.SubCategory1).
			OnConflict().
			UpdateNewValues().
			ID(c.Context)

		if err != nil {
			return err
		}

		create.SetSubCategory1ID(sub_category1_id)
	}

	if doc.SubCategory2 != "" {
		sub_category2_id, err := c.Client.FnetSubCategory2.Create().
			SetName(doc.SubCategory2).
			OnConflict().
			UpdateNewValues().
			ID(c.Context)

		if err != nil {
			return err
		}

		create.SetSubCategory2ID(sub_category2_id)
	}

	return create.OnConflict().
		UpdateNewValues().
		Exec(c.Context)
}
