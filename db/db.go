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

	return c.Client.FnetDocument.Create().
		SetFnetID(doc.ID).
		SetAdditionalInformation(doc.AdditionalInformation).
		SetDocumentCategory(doc.DocumentCategory).
		SetDocumentStatus(doc.DocumentStatus).
		SetDocumentSubCategory1(doc.DocumentSubCategory1).
		SetDocumentSubCategory2(doc.DocumentSubCategory2).
		SetFundDescription(doc.FundDescription).
		SetHighPriority(doc.HighPriority).
		SetMarketName(doc.MarketName).
		SetReferenceDate(referenceDate).
		SetReferenceDateFormat(doc.ReferenceDateFormat).
		SetReferenceDateStr(doc.ReferenceDate).
		SetReviewed(doc.Reviewed).
		SetStatus(doc.Status).
		SetStatusDescription(doc.StatusDescription).
		SetSubmissionDate(submissionDate).
		SetSubmissionDateStr(doc.SubmissionDate).
		SetSubmissionMethod(doc.SubmissionMethod).
		SetSubmissionMethodDescription(doc.SubmissionMethodDescription).
		SetVersion(doc.Version).
		OnConflict().
		UpdateNewValues().
		Exec(c.Context)
}
