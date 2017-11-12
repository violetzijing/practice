#include <stdio.h>

int func(int *arr) {
	printf("%ld\n", sizeof(arr));
}

int func2() {
	printf("baka\n");
	return 233;
}

int main(int argc, char *argv[]) {
	for(int i = 1; i < argc; i++) {
		printf("%s\n", argv[i]);
	}

	func2(2);
	int arr[87];
	int arr2[] = {3};

	func(arr);
	int a;
	char b;
	long c;
	short d;
	long long e;
	void *f, *g;

	printf("f: %ld\n", sizeof(f));
	printf("short: %ld\n", sizeof(short));
	printf("long: %ld\n", sizeof(long));
	printf("long long: %ld\n", sizeof(long long));

	printf("arr: %ld\n", sizeof(arr) / sizeof(int));
	printf("arr2: %ld\n", sizeof(arr2));
	return 0;
}


