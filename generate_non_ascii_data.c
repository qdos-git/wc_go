#include <stdio.h>

int main() {

  FILE *file;

  int i;

  file = fopen("output.bin", "wb");

  for ( i = 0 ; i < 8 ; i++ ) {

    fwrite(&i, sizeof(char), 1, file);

  }

  for ( i = 14 ; i < 32 ; i++ ) {

    fwrite(&i, sizeof(char), 1, file);

  }

}
