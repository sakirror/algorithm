#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <libgen.h>
#include "mysort.h"

static void printlist(const int *list, const int n)
{
    for (int i = 0; i < n; i++)
        printf("%d ", list[i]);
    printf("\n");
}

static int *preparelist(char *data[], int n)
{
    int *list = NULL;

    if ((list = (int *)malloc((n) * sizeof(int))) == NULL) {
        printf("malloc %d*%lu\n!!", n, sizeof(int));
        return NULL;
    }

    for (int i = 0; i < n; i++) {
        list[i] = atoi(data[i]);
    }

    return list;
}

int main(int argc, char *argv[])
{
    const struct {
        char *name;
        void (*sort)(int *, const int);
    } tbl[] = {
        {"insertion", insertion_sort},
    };
    int *list = NULL;

    if (argc < 2) {
        printf("need list!!\n");
        exit(1);
    }

    for (int i = 0; i < sizeof(tbl) / sizeof(tbl[0]); i++) {
        if (!strcmp(basename(argv[0]), tbl[i].name)) {
            int num = argc - 1;
            if (NULL == (list = preparelist(&argv[1], num))) {
                printf("error: preparelist\n");
                exit(1);
            }
            tbl[i].sort(list, num);
            printlist(list, num);
            break;
        }
    }

    if (list)
        free(list);
    else
        printf("invalid name:%s", basename(argv[0]));

    return 0;
}
