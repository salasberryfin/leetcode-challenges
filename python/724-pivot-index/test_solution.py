from solution import Solution
sol = Solution()

def test_example():
    input = [1,7,3,6,5,6]
    expected = 3
    result = sol.pivotIndex(input)

    assert result == expected


def test_example2():
    input = [1,2,3]
    expected = -1
    result = sol.pivotIndex(input)

    assert result == expected


def test_example3():
    input = [2,1,-1]
    expected = 0
    result = sol.pivotIndex(input)

    assert result == expected
