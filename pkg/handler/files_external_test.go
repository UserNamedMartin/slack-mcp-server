package handler

import (
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestIsSlackFileHost(t *testing.T) {
	assert.True(t, isSlackFileHost("https://files.slack.com/files-pri/T1-F1/download/a.pdf"))
	assert.True(t, isSlackFileHost("https://slack.com/x"))
	assert.False(t, isSlackFileHost("https://docs.google.com/presentation/d/abc/edit"))
	assert.False(t, isSlackFileHost("https://files.slack.com.evil.example/x"))
	assert.False(t, isSlackFileHost("http://files.slack.com/x"))
	assert.False(t, isSlackFileHost("::bad"))
}

func TestExtractFileID(t *testing.T) {
	assert.Equal(t, "F0000000001", extractFileID("F0000000001"))
	assert.Equal(t, "F0000000001", extractFileID(" F0000000001 "))
	// A team ID containing an "F" must not be mistaken for the file ID.
	assert.Equal(t, "F0000000001", extractFileID("https://files.slack.com/files-pri/T0FAKETEAM-F0000000001/download/a.pdf"))
	assert.Equal(t, "F0000000001", extractFileID("https://example.slack.com/files/U0000000001/F0000000001/a.pdf"))
}

func TestExternalFileResultDrive(t *testing.T) {
	f := &slack.File{
		ID:           "F0000000001",
		Name:         "example-presentation",
		Mimetype:     "application/vnd.google-apps.presentation",
		IsExternal:   true,
		ExternalType: "gdrive",
		URLPrivate:   "https://docs.google.com/presentation/d/EXAMPLE_drive_file_id_0123456789/edit?usp=sharing",
	}
	r := externalFileResult(f)
	assert.True(t, r.Skipped)
	assert.Empty(t, r.SavedPath)
	assert.Equal(t, "EXAMPLE_drive_file_id_0123456789", r.DriveFileID)
	assert.Equal(t, f.URLPrivate, r.ExternalURL)
}

func TestExternalFileResultOther(t *testing.T) {
	r := externalFileResult(&slack.File{ID: "F1", IsExternal: true, ExternalType: "dropbox", URLPrivate: "https://www.dropbox.com/s/xyz/a.pdf"})
	assert.True(t, r.Skipped)
	assert.Empty(t, r.DriveFileID)
	assert.Contains(t, r.Reason, "dropbox")
}
