class Solution():
    def isSubsequence(self, s: str, t: str) -> bool:
        matching = "".join("".join(x) for x in t if x in s)
        # print(f"Matching {matching}, s: {s} are equal: {matching == s}")
        
        """
        Two pointer:
            1. check if item from potential subsquence "s" matches item from
            matching characters found in "t"
            2. Keeps order
            3. If the iterator "i" is len(s): an ordered match has been
            found for all characters in "s"
        """

        if s == matching:
            return True

        i, j = 0, 0
        while (i <= len(s) - 1) and (j <= len(matching) - 1):
            if s[i] == matching[j]:
                i += 1
            if i == len(s):
                return True
            j += 1

        return False
