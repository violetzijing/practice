#include <stdio.h>
#include <string.h>

void my_strstr(char *str, char *substr) {
	int i;
	int str_length = strlen(str);
	int substr_length = strlen(substr);

	if (substr_length > str_length) {
		printf("nothing found");
		return;
	}
	for (i = 0; i <= (str_length - substr_length); i++) {
		if (str[i] != substr[0]) {
			continue;
		}
		// enter substr loop
		int j = 0;
		while (j < substr_length) {
			if (substr[j] != str[i+j]) {
				break;
			}

			j++;
		}

		if (j == substr_length) {
			printf("index: %d\n", i);
		}
	}
}

int main(int argc, char *argv[]) {
	my_strstr(argv[1], argv[2]);
//	my_strstr("001234123", "123");

	return 0;
}
