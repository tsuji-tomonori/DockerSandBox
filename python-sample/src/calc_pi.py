from __future__ import annotations


def calc_b2n(an: float, bn: float) -> float:
    return ((1/an) + (1/bn))**-1 * 2


def calc_a2n(an: float, b2n: float) -> float:
    return (an*b2n)**0.5


def calc(an: float, bn: float) -> tuple(float, float):
    b2n = calc_b2n(an, bn)
    a2n = calc_a2n(an, b2n)
    return (a2n, b2n)


if __name__ == "__main__":
    n = 6
    an = 3.0
    bn = 2*(3**0.5)
    while n < 10**8:
        an, bn = calc(an, bn)
        n *= 2
        print(f"n={n:010}, an={an:.50f}, bn={bn:.50f}")
    # 3.14159265358979