from solution import Solution
sol = Solution()

def test_pattern_example1():
    pattern = "abba"
    s = "dog cat cat dog"

    result = sol.wordPattern(pattern, s)
    expected = True

    assert result == expected


def test_pattern_example2():
    pattern = "abba"
    s = "dog cat cat fish"

    result = sol.wordPattern(pattern, s)
    expected = False

    assert result == expected


def test_pattern_example3():
    pattern = "aaaa"
    s = "dog cat cat dog"

    result = sol.wordPattern(pattern, s)
    expected = False

    assert result == expected


def test_pattern_example4():
    pattern = "abba"
    s = "dog dog dog dog"

    result = sol.wordPattern(pattern, s)
    expected = False

    assert result == expected


def test_pattern_example5():
    pattern = "abc"
    s = "b c a"

    result = sol.wordPattern(pattern, s)
    expected = True

    assert result == expected


def test_pattern_example6():
    pattern = "aba"
    s = "cat cat cat dog"

    result = sol.wordPattern(pattern, s)
    expected = False

    assert result == expected


def test_pattern_example7():
    pattern = "aba"
    s = "dog cat cat"

    result = sol.wordPattern(pattern, s)
    expected = False

    assert result == expected
