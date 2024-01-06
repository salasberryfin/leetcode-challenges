from solution import Solution
sol = Solution()


def test_example1():
    startTime = [1,2,3,3]
    endTime = [3,4,5,6]
    profit = [50,10,40,70]

    profit = sol.jobScheduling(startTime, endTime, profit)
    expected = 120

    assert profit == expected


def test_example2():
    startTime = [1,2,3,4,6]
    endTime = [3,5,10,6,9]
    profit = [20,20,100,70,60]

    profit = sol.jobScheduling(startTime, endTime, profit)
    expected = 150

    assert profit == expected


def test_example3():
    startTime = [1,1,1]
    endTime = [2,3,4]
    profit = [5,6,4]

    profit = sol.jobScheduling(startTime, endTime, profit)
    expected = 6

    assert profit == expected
