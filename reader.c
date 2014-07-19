#include <archive.h>
#include <archive_entry.h>
#include "_cgo_export.h"

ssize_t go_libarchive_reader(struct archive *a, void *client_data, const void **block) {
	struct archiveReader_return ret = archiveReader(a, client_data);
	*block = ret.r0;
	ssize_t size = (ssize_t) ret.r1;

	return size;
}

void go_libarchive_set(struct archive *a) {
	archive_read_set_read_callback(a, go_libarchive_reader);
}

struct go_libarchive_read_next_ret* go_libarchive_read_next(struct archive *a) {
	struct archive_entry *e;
	int result = archive_read_next_header(a, &e);

	struct go_libarchive_read_next_ret *ret = (struct go_libarchive_read_next_ret *) malloc(sizeof(struct go_libarchive_read_next_ret));

	ret->e = e;
	ret->result = result;

	return ret;
}