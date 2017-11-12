#include <stdio.h>
#include <stdlib.h>

int *merge(int arr1[], int n1, int arr2[], int n2);

int main(int argc, char *argv[]) {
	int arr1[] = {1, 4, 5, 6};
	int n1 = sizeof(arr1) / sizeof(int);

	int arr2[] = {2, 2, 5, 7, 8};
	int n2 = sizeof(arr2) / sizeof(int);

	int *arr = merge(arr1, n1, arr2, n2);

	for (int i = 0; i < n1 + n2; i++) {
		printf("%d, ", arr[i]);
	}
	return 0;
}

int *merge(int arr1[], int n1, int arr2[], int n2) {
	int *arr3 = (int *)malloc(sizeof(int) * (n1 + n2));

	int i = 0, j = 0, k = 0;

	while (i != n1 && j != n2) {
		if (arr1[i] > arr2[j]) {
			arr3[k++] = arr2[j];
			j++;
		} else {
			arr3[k++] = arr1[i];
			i++;
		}
	}

	if (i != n1) {
		while(i != n1) {
			arr3[k++] = arr1[i++];
		}
	} else {
		while(j != n2) {
			arr3[k++] = arr2[j++];
		}
	}

	return arr3;
}
