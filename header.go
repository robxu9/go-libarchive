package archive

/*
#cgo pkg-config: libarchive
#include <archive.h>
#include <archive_entry.h>
*/
import "C"
import "time"

type Header struct {
	AccessTime time.Time // access time
	BirthTime  time.Time // creation time
	ChangeTime time.Time // status change time

	DevCombined int64 // combined # of charcter or block device
	DevMajor    int64 // major # of character or block device
	DevMinor    int64 // minor # of character or block device

	FileType int64 // file type this header represents
}
