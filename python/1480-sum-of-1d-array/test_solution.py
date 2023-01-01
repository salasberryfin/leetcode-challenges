from solution import Solution
sol = Solution()


def test_example1():
    nums = [1, 2, 3, 4]

    result = sol.runningSum(nums)
    expected = [1, 3, 6, 10]

    assert result == expected


def test_example2():
    nums = [1,1,1,1,1]

    result = sol.runningSum(nums)
    expected = [1, 2, 3, 4, 5]

    assert result == expected


def test_example3():
    nums = [3,1,2,10,1]

    result = sol.runningSum(nums)
    expected = [3,4,6,16,17]

    assert result == expected
