#define _XOPEN_SOURCE 500

#include <ftw.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <stdio.h>
#include <sys/inotify.h>
#include <limits.h>
#include <unistd.h>
#include "logger.h"

#define BUF_LEN (10 * (sizeof(struct inotify_event) + NAME_MAX + 1))

//Alejandro Castro Arévalo-A01636235

int inotifyFd;

static int get_files(const char *fpath, const struct stat *sb, int tflag, struct FTW *ftwbuf) {
	int wd = inotify_add_watch(inotifyFd, fpath, IN_CREATE | IN_DELETE | IN_MOVED_FROM | IN_MOVED_TO);
	if (wd == -1) {
		errorf("Add to inotify_add_watch Failed");
		exit(1);
	}
	return 0;
}

static void displayInotifyEvent(struct inotify_event *i) {
	if (i->mask & IN_CREATE){
        if(i->mask & IN_ISDIR){
            infof("[CREATE-DIRECTORY] -");
        }else{
            infof("[CREATE-FILE] -");
        }
    }	
	if (i->mask & IN_DELETE){
        if(i->mask & IN_ISDIR){
            infof("[DELETE-DIRECTORY] -");
        }else{
            infof("[DELETE-FILE] -");
        }	
    }
	if (i->mask & IN_MOVED_FROM){
        if(i->mask & IN_ISDIR){
            infof("[RENAME DIRECTORY-FROM] -");
        }else{
            infof("[RENAME FILE-FROM] -");
        }
    }		
	if (i->mask & IN_MOVED_TO){
        if(i->mask & IN_ISDIR){
            infof("[RENAME DIRECTORY-TO] -");
        }else{
            infof("[RENAME FILE-TO] -");
        }
    }

	if (i->len > 0)
		printf(" %s\n", i->name);
}

int main(int argc, char *argv[]) {
	inotifyFd = inotify_init();
	if (inotifyFd == -1) {
		errorf("Create inotifyFd Failed");
		exit(1);
	}

	int flags = FTW_PHYS;
	if (nftw((argc < 2) ? "." : argv[1], get_files, 20, flags) == -1) {
		panicf("Traverse nftw Failed");
		exit(1);
	}

	char buf[BUF_LEN] __attribute__ ((aligned(8)));
	ssize_t numRead;
	char *p;
	struct inotify_event *event;

	for (;;) {		/*Indefinite loop*/
		numRead = read(inotifyFd, buf, BUF_LEN);
		if (numRead == 0) {
			panicf("read() from inotify fd returned 0!");
			exit(1);
		}
		if (numRead == -1) {
			errorf("read");
			exit(1);
		}
		for (p = buf; p < buf + numRead;) {
			event = (struct inotify_event *)p;
			displayInotifyEvent(event);
			p += sizeof(struct inotify_event) + event->len;
		}
		inotifyFd = inotify_init();
		if (nftw((argc < 2) ? "." : argv[1], get_files, 20, flags) == -1) {
			panicf("Transverse nftw Failed");
			exit(1);
		}
	}
	exit(0);
}