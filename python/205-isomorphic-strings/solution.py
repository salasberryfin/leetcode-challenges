class Solution:
    """
    Can be improved using zip(s, t) and generating tuples containing interlocked
    values from both s and t
    """

    def isIsomorphic(self, s: str, t: str) -> bool:
        """
        1. Get positions in the strings for each character: List[List]
        2. Check if each list of positions exist in both lists
        3. To improve performance, check positions while iterating over
        second list
        """
        if len(s) != len(t):
            return False

        # using a set avoids going multiple times over repeated values
        s_set = set(s)
        t_set = set(t)

        appearances_s = []
        for item in s_set:
            positions = [x for x in range(len(s)) if s[x] == item]
            appearances_s.append(positions)

        for item in t_set:
            positions = [x for x in range(len(t)) if t[x] == item]
            if positions not in appearances_s:
                return False

        # print(f"Appearances S: {appearances_s}, appearances T: {appearances_t}")

        return True
