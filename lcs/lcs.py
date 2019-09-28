# -*- coding: utf-8 -*-
import sys


def print_lcs (a, b, lcs, i, j):
    if (i == 0 or j == 0):
        return
    if (a[i-1] == b[j-1]):
        print_lcs(a, b, lcs, i - 1, j - 1)
        print(a[i - 1], end="")
    else:
        if lcs[i-1][j] >= lcs[i][j-1]:
            print_lcs(a, b, lcs, i - 1, j)
        else:
            print_lcs(a, b, lcs, i, j - 1)

if __name__ == '__main__' and len(sys.argv) > 2:
    s1 = sys.argv[1]
    s2 = sys.argv[2]
    lcs = [[0 for j in range(len(s2)+1)] for i in range(len(s1)+1)]
    for i in range(1, len(s1)+1):
        for j in range(1, len(s2)+1):
            a = 1 if s1[i-1]==s2[j-1] else 0
            lcs[i][j] = max(lcs[i-1][j-1] + a, lcs[i-1][j], lcs[i][j-1])

    for row in lcs:
        print(row)
    print("lcs is", lcs[len(s1)][len(s2)])
    print_lcs(s1, s2, lcs, len(s1), len(s2))
    print("")
