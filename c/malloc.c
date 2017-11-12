#include <stdio.h>
#include <stdlib.h>

void func233(int);
#define laowang 4786

#define multi(a, b)		((a) * (b))

int main(int argc, char *argv[]) {
	int *baka = (int*)malloc(23 * sizeof(int));
	baka[0] = 233;
	baka[1] = 2333;

	for(int i = 0; i < 23; i++) {
		printf("num: %d\n", baka[i]);
	}

	func233(2);

	laowang+1;

	const int lobster = 99;
	lobster = 45;
	printf("lobster: %d\n", lobster);
	return 0;
}

void func233(int x) {
	printf("%d\n", ++x);
}
