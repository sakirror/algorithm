#!/usr/bin/env python
# -*- coding: utf-8 -*-

import sys

if __name__ == '__main__' and len(sys.argv) > 2:
    s1 = sys.argv[1]
    s2 = sys.argv[2]

    m = [[0] * (len(s2)+1) for i in range(len(s1)+1)]
    for i in range(len(s1)+1):
        m[i][0] = i
    for j in range(len(s2)+1):
        m[0][j] = j

    for i in range(1, len(s1)+1):
        for j in range(1, len(s2)+1):
            if s1[i-1] == s2[j-1]:
                x = 0
            else:
                x = 1
            m[i][j] = min(m[i-1][j] + 1, m[i][j-1] + 1, m[i-1][j-1] + x)

    print(" ", [x for x in s2])
    s1 = " " + s1
    for i,row in enumerate(m):
        print(s1[i], row)

    print("\n", m[-1][-1])
