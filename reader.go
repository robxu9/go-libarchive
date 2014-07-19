package archive

//#cgo pkg-config: libarchive
//#include <archive.h>
//#include <archive_entry.h>
//#include <stdlib.h>
//#include "reader.h"
import "C"
import (
	"io"
	"os"
	"runtime"
	"time"
	"unsafe"
)

/*-
 * Basic outline for reading an archive:
 *   1) Ask archive_read_new for an archive reader object.
 *   2) Update any global properties as appropriate.
 *      In particular, you'll certainly want to call appropriate
 *      archive_read_support_XXX functions.
 *   3) Call archive_read_open_XXX to open the archive
 *   4) Repeatedly call archive_read_next_header to get information about
 *      successive archive entries.  Call archive_read_data to extract
 *      data for entries of interest.
 *   5) Call archive_read_finish to end processing.
 */

type Reader struct {
	archive *C.struct_archive
	reader  io.Reader
}

func NewReader(r io.Reader) *Reader {
	reader := &Reader{}
	reader.archive = C.archive_read_new()
	C.archive_read_support_filter_all(reader.archive)
	C.archive_read_support_format_all(reader.archive)

	reader.reader = r

	C.go_libarchive_set(reader.archive)
	C.archive_read_set_callback_data(reader.archive, unsafe.Pointer(reader))

	runtime.SetFinalizer(reader, func(r *Reader) {
		C.archive_read_free(r.archive)
	})

	return reader
}

//export archiveReader
func archiveReader(a *C.struct_archive, client_data unsafe.Pointer) (unsafe.Pointer, int) {
	reader := (*Reader)(client_data)
	buffer := make([]byte, 1024)
	read, err := reader.reader.Read(buffer)

	ret := unsafe.Pointer(&buffer[0])
	if read == 0 && err == io.EOF {
		return ret, 0
	} else if err != io.EOF {
		return ret, ARCHIVE_FATAL
	}
	return ret, read
}

func (r *Reader) Next() (*Header, error) {
	result := C.go_libarchive_read_next(r.archive)
	defer C.free(unsafe.Pointer(result))

	switch result.result {
	case ARCHIVE_EOF:
		return nil, ErrArchiveEOF
	case ARCHIVE_FATAL:
		return nil, ErrArchiveFatal
	case ARCHIVE_RETRY:
		return nil, ErrArchiveRetry
	case ARCHIVE_WARN:
		return nil, ErrArchiveWarn
	}

	header := &Header{}

	if C.archive_entry_atime_is_set(result.e) != 0 {
		sec := C.archive_entry_atime(result.e)
		nsec := C.archive_entry_atime_nsec(result.e)
		header.AccessTime = time.Unix(int64(sec), int64(nsec))
	}

	if C.archive_entry_birthtime_is_set(result.e) != 0 {
		sec := C.archive_entry_birthtime(result.e)
		nsec := C.archive_entry_birthtime_nsec(result.e)
		header.BirthTime = time.Unix(int64(sec), int64(nsec))
	}

	if C.archive_entry_ctime_is_set(result.e) != 0 {
		sec := C.archive_entry_ctime(result.e)
		nsec := C.archive_entry_ctime_nsec(result.e)
		header.ChangeTime = time.Unix(int64(sec), int64(nsec))
	}

	if C.archive_entry_dev_is_set(result.e) != 0 {
		header.DevCombined = int64(C.archive_entry_dev(result.e))
		header.DevMajor = int64(C.archive_entry_devmajor(result.e))
		header.DevMinor = int64(C.archive_entry_devminor(result.e))
	}

	header.FileType = int64(C.archive_entry_filetype(result.e))

	return header, nil
}

// When completed, will return (0, io.EOF)
func (r *Reader) Read(b []byte) (n int, err error) {
	return 0, os.ErrNotExist
}
