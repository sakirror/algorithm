#include <stdio.h>
#include <string.h>

static void swap(int *list, const int i1, const int i2)
{
    int tmp = list[i1];
    list[i1] = list[i2];
    list[i2] = tmp;
}

static void printh(int *list, int i)
{
    printf("%d:%d ", i, list[i]);
}

static void move(int *list, const int index, const int max)
{
    int tmp = index;
    int left = tmp * 2 + 1;
    int right = left + 1;

    if (left < max && list[left] > list[tmp])
        tmp = left;
    if (right < max && list[right] > list[tmp])
        tmp = right;

    if (tmp != index) {
        swap(list, index, tmp);
        move(list, tmp, max);
    }
}

void heap_sort(int *list, const int n)
{
    for (int i = n / 2 - 1; i >= 0; i--)
        move(list, i, n);

    for (int i = n - 1; i > 0; i--) {
        swap(list, 0, i);
        move(list, 0, i);
    }
}
