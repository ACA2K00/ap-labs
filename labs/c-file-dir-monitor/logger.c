#include <stdio.h>
#include "logger.h"

#include <stdarg.h>
#include <stdlib.h>
#include <syslog.h>
#include <string.h>

//Alejandro Castro Arévalo-A01636235

int loggerType;

int printFormat(const char *format, va_list arg){
  int res = 0;
  if(loggerType){
      vsyslog(LOG_INFO, format, arg);
  }else{
      res = vprintf(format, arg);
  }
	va_end(arg);
    //printf("\n");
	return res;
}

int initLogger(char *logType) {

  if(!logType || strcmp(logType, "stdout") == 0 || !logType[0]){
    if(loggerType == 1){
      printf("Initializing Logger on: stdout\n");
      loggerType = 0;
      closelog();
    } else{
      printf("Already initialized on: stdout\n");
    }
  }
  else if(strcmp(logType, "syslog") == 0){
    if(loggerType == 1){
      printf("Initializing Logger on: syslog\n");

      openlog(NULL, LOG_NDELAY, LOG_USER);
      loggerType = 0;
    } else{
      printf("Already initialized on: syslog\n");
    }
  }
  else{
    errorf("Logger option '%s' is unknown\n", logType);
    return -1;
  }
  return 0;
}


int infof(const char *format, ...){
	va_list arg;
	va_start(arg, format);
	return printFormat(format, arg);
}

int warnf(const char *format, ...){
	va_list arg;
	va_start(arg, format);
	return printFormat(format, arg);
}

int errorf(const char *format, ...){
	va_list arg;
	va_start(arg, format);
	return printFormat(format, arg);
}

int panicf(const char *format, ...){
	va_list arg;
	va_start(arg, format);
	printFormat(format, arg);
	abort();
	return -1;
}