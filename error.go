package archive

import (
	"errors"
)

const (
	ARCHIVE_EOF   = 1
	ARCHIVE_OK    = 0
	ARCHIVE_RETRY = -10
	ARCHIVE_WARN  = -20
	ARCHIVE_FATAL = -30
)

var (
	ErrArchiveEOF   = errors.New("libarchive: EOF [end of file]")
	ErrArchiveRetry = errors.New("libarchive: RETRY [operation failed but can be retried]")
	ErrArchiveWarn  = errors.New("libarchive: WARN [success but non-critical error]")
	ErrArchiveFatal = errors.New("libarchive: FATAL [critical error, archive closing]")
)
