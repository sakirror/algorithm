# -*- coding: utf-8 -*-

import math
import random
import matplotlib.pyplot as plt
import numpy as np

class Point:
    """
    座標用
    """
    def __init__(self, x, y):
        self.x = x
        self.y = y

class Graham:
    """
    グラハム走査
    """
    def __init__(self, num, xlim, ylim):
        point = lambda i : Point(random.randrange(xlim), random.randrange(ylim))
        # ソート済点集合
        self._points = self._sort([point(i) for i in range(num)])

    def _sort(self, points):
        """
        最下点を基準にした角度ソート済点集合. p[0]:最下点, p[1 ~ -1]:角度ソート
        """
        low = points[0]
        for p in points:
            if p.y < low.y or (p.y == low.y and  p.x < low.x):
                low = p
        points.remove(low)
        if len(points) < 3:
            return [low, points[0]]

        # 角度でソートする.同じ角度の場合は大きさで
        f_sort = lambda p: (math.atan2(p.y - low.y, p.x - low.x), (p.y - low.y)**2 + (p.x - low.x)**2)
        points.sort(key=f_sort)

        # 重複は集合に含めない. sort済なので直前の値とのみ比較
        return [low, points[0]] + \
            [points[i] for i in range(1, len(points)) if points[i].x != points[i-1].x or points[i].y != points[i-1].y]

    def get_points(self):
        return self._points

    def search(self, points):
        # 3点までなら確実に凸包
        if len(points) <= 3:
            return points

        edge = [points[i] for i in range(3)]
        for i in range(3, len(points)):
            while self._is_clockwise(edge[-2], edge[-1], points[i]) is True and len(edge) >= 3:
                edge.pop(-1)
            edge.append(points[i])
        return edge

    def _is_clockwise(self, p1, p2, p3):
        """
        外積を使って時計回りか判定する.
        """
        return (p2.y - p1.y)*(p3.x - p1.x) - (p2.x - p1.x)*(p3.y - p1.y) > 0

    def _get_xylist(self, l):
        x = list()
        y = list()
        for p in l:
            x.append(p.x)
            y.append(p.y)
        return x, y

    def print_points(self, points):
        for p in points:
            print("[{}, {}],".format(p.x, p.y), end='')
        print()

    def draw(self, points, convex=None):
        """
        凸包がわたされている場合は線を引く.
        """
        fig, ax = plt.subplots(1,1)
        x, y = self._get_xylist(points)
        ax.plot(x, y, 'o')
        if convex:
            # 凸包の線を書くときに終端を最下点にしておかないと1周しない
            x, y = self._get_xylist(convex + [convex[0]])
            ax.plot(x, y)
        plt.show()

def main():
    for i in range(3):
        graham = Graham(random.randrange(1,50), random.randrange(1,30) , random.randrange(1,30))
        points = graham.get_points()
        edges = graham.search(points)
        # graham.print_points(points)
        # graham.print_points(edges)
        graham.draw(points, edges)


if __name__ == '__main__':
    main()
