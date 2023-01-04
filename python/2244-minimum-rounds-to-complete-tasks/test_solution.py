from solution import Solution

sol = Solution()


def test_example1():
    tasks = [2, 2, 3, 3, 2, 4, 4, 4, 4, 4]
    expected = 4
    result = sol.minimumRounds(tasks)

    assert  expected == result


def test_example2():
    tasks = [2, 3, 3]
    expected = -1
    result = sol.minimumRounds(tasks)

    assert  expected == result


def test_example3():
    tasks = [7,7,7,7,7,7]
    expected = 2
    result = sol.minimumRounds(tasks)

    assert  expected == result


def test_example4():
    tasks = [69,65,62,64,70,68,69,67,60,65,69,62,65,65,61,66,68,61,65,63,60,66,68,66,67,65,63,65,70,69,70,62,68,70,60,68,65,61,64,65,63,62,62,62,67,62,62,61,66,69]
    expected = 20
    result = sol.minimumRounds(tasks)

    assert  expected == result
