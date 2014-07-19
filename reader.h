#include <archive.h>
#include <archive_entry.h>

ssize_t go_libarchive_reader(struct archive *a, void *client_data, const void **block);

void go_libarchive_set(struct archive *a);

struct go_libarchive_read_next_ret {
	struct archive_entry *e;
	int result;
};

struct go_libarchive_read_next_ret* go_libarchive_read_next(struct archive *a);