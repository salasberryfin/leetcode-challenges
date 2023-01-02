from solution import Solution

sol = Solution()


def test_example1():
    s = "egg"
    t = "add"
    expected = True
    result = sol.isIsomorphic(s, t)

    assert result == expected


def test_example2():
    s = "foo"
    t = "bar"
    expected = False
    result = sol.isIsomorphic(s, t)

    assert result == expected


def test_example3():
    s = "paper"
    t = "title"
    expected = True
    result = sol.isIsomorphic(s, t)

    assert result == expected


def test_example4():
    s = "badc"
    t = "baba"
    expected = False
    result = sol.isIsomorphic(s, t)

    assert result == expected


def test_example5():
    s = "bbbaaaba"
    t = "aaabbbba"
    expected = False
    result = sol.isIsomorphic(s, t)

    assert result == expected
