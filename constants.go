package archive

//#cgo pkg-config: libarchive
//#include "constants.h"
import "C"

var (
	FileTypeRegFile = C.go_libarchive__AE_IFREG()
	FileTypeSymLink = C.go_libarchive__AE_IFLNK()
	FileTypeSocket  = C.go_libarchive__AE_IFSOCK()
	FileTypeCharDev = C.go_libarchive__AE_IFCHR()
	FileTypeBlkDev  = C.go_libarchive__AE_IFBLK()
	FileTypeDir     = C.go_libarchive__AE_IFDIR()
	FileTypeFIFO    = C.go_libarchive__AE_IFIFO()
)
