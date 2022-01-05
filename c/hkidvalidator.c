#include "hkidvalidator.h"

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

void slUpper(char *s, int len);
int slIsDigits(char *s, int len);
int hkidSum(char *hkid);
int checkDigit(char *hkid);

int checkHKID(char *hkid) {
  int l = strlen(hkid);
  char s[9];

  switch (l) {
  case 8:
    s[0] = ' ';
    char *ss = &s[1];
    strcpy(ss, hkid);
    break;
  case 9:
    if (isalpha(s[0])) {
      break;
    }
    /* If not, fall through. */
  default:
    return 0;
  }
  if (isalpha(s[1]) && slIsDigits(&s[2], 6) && (s[8] == 'A' || isdigit(s[8]))) {
    slUpper(s, 9);
    int sum = hkidSum(s);
    int cd = checkDigit(s);
    int mo = sum % 11;

    printf("sum:%d, cd:%d, mo:%d\n", sum, cd, mo);
    if (mo == 0) {
      return cd == 0;
    } else {
      return cd == 11 - mo;
    }
  } else {
    return 0;
  }
}

void slUpper(char *s, int len) {
  for (int i = 0; i < len; i++) {
    s[i] = toupper(s[i]);
  }
}

int slIsDigits(char *s, int len) {
  for (int i = 0; i < len; i++) {
    if (!isdigit(s[i])) {
      return 0;
    }
  }
  return 1;
}

int hkidSum(char *hkid) {
  /* Length must larger than 8. */
  int sum = 0;
  for (int i = 0; i < 8; i++) {
    if (hkid[i] == ' ') {
      sum += 36 * 9;
    } else if (isalpha(hkid[i])) {
      sum += ((int)hkid[i] - 'A' + 10) * (9 - i);
    } else {
      sum += ((int)hkid[i] - '0') * (9 - i);
    }
  }
  return sum;
}

int checkDigit(char *hkid) {
  if (hkid[8] == 'A') {
    return 10;
  } else {
    return hkid[8] - '0';
  }
}
