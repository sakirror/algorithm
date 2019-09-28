#include <stdio.h>
#include <string.h>

inline static int getmax(int a, int b, int c)
{
    return (a >= b && a >= c) ? a : ((b >= c) ? b : c);
}

static void printlcstable(char (*lcs)[100], int imax, int jmax)
{
    for (int i = 0; i < imax; i++) {
        for (int j = 0; j < jmax; j++) {
            printf("%02d ", lcs[i][j]);
        }
        printf("\n");
    }
}

static void printlcsword(char (*lcs)[100], char *s1, char *s2, int i, int j)
{
    if (i == 0 || j == 0)
        return;
    if (s1[i - 1] == s2[j - 1]) {
        printlcsword(lcs, s1, s2, i - 1, j - 1);
        printf("%c[%d]", s1[i - 1], i);
    } else if (lcs[i - 1][j] >= lcs[i][j - 1]) {
        printlcsword(lcs, s1, s2, i - 1, j);
    } else {
        printlcsword(lcs, s1, s2, i, j - 1);
    }
}

int main(int argc, char **argv)
{
    char *s[2];
    int slen[2] = {};
    char lcs[100][100] = {0};

    if (argc < 3 || strlen(argv[1]) >= 100 || strlen(argv[2]) >= 100) {
        printf("invalid argc %d\n", argc);
        return 1;
    }

    for (int i = 0; i < 2; i++) {
        s[i] = argv[i + 1];
        slen[i] = strlen(s[i]);
    }

    for (int i = 1; i < slen[0] + 1; i++) {
        for (int j = 1; j < slen[1] + 1; j++) {
            lcs[i][j] = getmax(lcs[i - 1][j - 1] + (s[0][i - 1] == s[1][j - 1]) ? 1 : 0
                              ,lcs[i - 1][j]
                              ,lcs[i][j - 1]);
        }
    }

    printlcstable(lcs, slen[0] + 1, slen[1] + 1);
    printf("\n");
    printlcsword(lcs, s[0], s[1], slen[0] + 1, slen[1] + 1);
    printf("\n");
    return 0;
}
