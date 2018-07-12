#include <string.h>

void insertion_sort(int *list, const int n)
{
    for (int i = 1; i < n; i++) {
        int val = list[i];
        for (int j = 0; j < i; j++) {
            if (list[i] < list[j]) {
                memmove(&list[i], &list[i + 1], (n - i - 1) * sizeof(int));
                memmove(&list[j + 1], &list[j], (n - j - 1) * sizeof(int));
                list[j] = val;
                break;
            }
        }
    }
}
