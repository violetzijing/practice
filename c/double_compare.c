#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int my_cmp(const void *a, const void *b) {
	double *p1 = (double *)a;
	double *p2 = (double *)b;
	if (*p1 > *p2) {
		return 1;
	} else if (*p1 == *p2) {
		return 0;
	} else {
		return -1;
	}
}

int main(int argc, char *argv[]) {
	double arr1[] = {1, 4, 5, 7, 3, 7};
	qsort(arr1, sizeof(arr1) / sizeof(double), sizeof(double), my_cmp);
	for (int i = 0; i < (sizeof(arr1) / sizeof(double)); i++) {
		printf("%lf, ", arr1[i]);
	}
	printf("\n");

	int (*ptr)(const void *a, const void *b);

	ptr = &my_cmp;
	double a1 = 3;
	double b1 = 4;
	int test = (*ptr)(&a1, &b1);
	printf("%d\n", test);
	return 0;
}
