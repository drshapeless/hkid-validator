#include "hkidvalidator.h"

#include <stdio.h>

int main(int argc, char *argv[]) {
  char *id = argv[1];
  int isReal = checkHKID(id);
  if (isReal) {
    printf("Real\n");
  } else {
    printf("Fake\n");
  }
  return 0;
}
