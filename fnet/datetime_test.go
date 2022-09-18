package fnet_test

import (
	"testing"
	"time"

	"github.com/lenon/gofii/fnet"
	"github.com/stretchr/testify/assert"
)

func TestParseDateWithValidFormat(t *testing.T) {
	actual, err := fnet.ParseDate("11/2019", fnet.DATE_FORMAT_MY)
	assert.Nil(t, err)
	assert.Equal(t, time.Date(2019, 11, 1, 0, 0, 0, 0, time.Local), actual)

	actual, err = fnet.ParseDate("21/11/2019", fnet.DATE_FORMAT_DMY)
	assert.Nil(t, err)
	assert.Equal(t, time.Date(2019, 11, 21, 0, 0, 0, 0, time.Local), actual)

	actual, err = fnet.ParseDate("21/11/2019 16:20", fnet.DATE_FORMAT_DMY_HM)
	assert.Nil(t, err)
	assert.Equal(t, time.Date(2019, 11, 21, 16, 20, 0, 0, time.Local), actual)
}

func TestParseDateWithInvalidFormat(t *testing.T) {
	_, err := fnet.ParseDate("11/2019", fnet.DATE_FORMAT_DMY)
	assert.EqualError(t, err, "parsing time \"11/2019\": month out of range")

	_, err = fnet.ParseDate("21/11/2019", fnet.DATE_FORMAT_MY)
	assert.EqualError(t, err, "parsing time \"21/11/2019\": month out of range")

	_, err = fnet.ParseDate("21/11/2019 16:20", fnet.DATE_FORMAT_DMY)
	assert.EqualError(t, err, "parsing time \"21/11/2019 16:20\": extra text: \" 16:20\"")

	_, err = fnet.ParseDate("21/11/2019 16:20", "5")
	assert.EqualError(t, err, "unknown format: 5")
}
