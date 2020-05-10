#include "query.h"
#include <stdio.h>
#include <stdlib.h>

int query(const char *query, char ***respout) {
  int i, l;
	char **resp;

  l = atoi(query);
  if (l <= 0) {
    return 0;
  }

  resp = calloc(l, sizeof(char*));
  for (i = 0; i < l; i++) {
    resp[i] = calloc(6, sizeof(char));
    sprintf(resp[i], "r-%03d", i);
  }
	*respout = resp;
  return l;
}
