// Author: Andreas Fladstad
// Compile with gcc roller.c -o roller

#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main(int argc, char** argv) {
  if(argc != 2) {
    printf("Usage: ./%s <roll>", argv[0]);
    return -1;
  }

  // Parse dice

  int amount = 0;
  int dice = 0;
  int plus = 0;
  char* input = argv[1];
  char buffer[20];
  int count = 0;

  while(*input != 'd') {
    if(*input == 0) {
      printf("error in input\n");
    }
    buffer[count++] = *input;
    input++;
  }
  input++;
  buffer[count] = 0;

  if(count != 0) {
    amount = atoi(buffer);
  }

  count = 0;

  while(*input != '+' && *input != 0) {
    buffer[count++] = *input;
    input++;
  }

  buffer[count] = 0;
  dice = atoi(buffer);

  if(*input == '+') {
    count = 0;
    input++;
    while(*input != 0) {
      buffer[count++] = *input;
      input++;
    }
    buffer[count] = 0;
    plus = atoi(buffer);
  }

  // Roll dice
  srand(time(NULL));
  int result = 0;
  for(int i = 0; i < amount; i++) {
    int current = rand() % dice + 1;
    result += current;
    printf("%d ", current);
  }
  if(plus != 0) {
    printf("+ %d", plus);
  }
  printf(" = %d\n", result);
  return 0;
}
