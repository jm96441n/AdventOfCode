from __future__ import annotations
from dataclasses import dataclass


@dataclass
class Point:
    x: int
    y: int


@dataclass
class Vector:
    start: Point
    end: Point

    @classmethod
    def new_from_input(cls, input: str) -> Vector:
        points = input.split(" -> ")
        start = points[0].split(",")
        end = points[1].split(",")
        p1 = Point(x=int(start[0]), y=int(start[1]))
        p2 = Point(x=int(end[0]), y=int(end[1]))
        return Vector(start=p1, end=p2)

    def points(self):
        if self.is_vertical():
            x_change = 0
            y_change = 1
            x_pred = lambda cur_x, x: cur_x <= x
            y_pred = lambda cur_y, y: cur_y <= y

            if self.start.y > self.end.y:
                y_change = -1
                y_pred = lambda cur_y, y: cur_y >= y
        elif self.is_horizontal():
            x_change = 1
            y_change = 0
            x_pred = lambda cur_x, x: cur_x <= x
            y_pred = lambda cur_y, y: cur_y <= y
            if self.start.x > self.end.x:
                x_change = -1
                x_pred = lambda cur_x, x: cur_x >= x
        else:
            x_change = 1
            y_change = 1
            x_pred = lambda cur_x, x: cur_x <= x
            y_pred = lambda cur_y, y: cur_y <= y
            if self.start.x > self.end.x:
                x_change = -1
                x_pred = lambda cur_x, x: cur_x >= x
            if self.start.y > self.end.y:
                y_change = -1
                y_pred = lambda cur_y, y: cur_y >= y
        cur_x = self.start.x
        cur_y = self.start.y
        while (y_pred(cur_y, self.end.y) or self.is_horizontal()) and (
            x_pred(cur_x, self.end.x or self.is_vertical())
        ):
            yield (cur_x, cur_y)
            cur_x += x_change
            cur_y += y_change

    def is_vertical(self):
        return self.start.x == self.end.x

    def is_horizontal(self):
        return self.start.y == self.end.y


def parttwo(vectors: list[Vector]) -> int:
    points: dict[tuple[int, int], int] = {}
    for v in vectors:
        for p in v.points():
            points[(p[0], p[1])] = points.get((p[0], p[1]), 0) + 1

    danger_zones = 0
    for k, v in points.items():
        if v > 1:
            danger_zones += 1
    return danger_zones


def partone(vectors: list[Vector]) -> int:
    points: dict[tuple[int, int], int] = {}
    filtered_vectors = filter(
        (lambda v: (v.is_vertical() or v.is_horizontal())), vectors
    )
    for v in filtered_vectors:
        for p in v.points():
            points[(p[0], p[1])] = points.get((p[0], p[1]), 0) + 1

    danger_zones = 0
    for _, v in points.items():
        if v > 1:
            danger_zones += 1
    return danger_zones


def run():
    with open("./input.txt") as f:
        vectors = [Vector.new_from_input(l) for l in f.read().splitlines()]
    print(partone(vectors))
    print(parttwo(vectors))


if __name__ == "__main__":
    run()
