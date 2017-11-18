#include <stdio.h>
#include <stdlib.h>

int *compar(const void *)

void three_sum(int nums[], int n) {
  int m = 0;
  while (m < n - 3) {
    int i = m + 1;
    int j = n - 1;

    while (i < j) {
      int result = nums[m] + nums[i] + nums[j];
      if (result == 0) {
        printf("%d = %d + %d + %d", result, nums[m], nums[i], nums[j]);
        i++;
        j--;
      } else if (result > 0) {
        j--;
      } else {
        i++;
      }

      m++;
    }
  }
}
int main(int argc, char *argv[]) {
  int nums[] = {-2, 0, 2}
  three_sum()
  return 0;
}

