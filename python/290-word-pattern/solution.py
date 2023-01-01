class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        s_list = s.split()

        if len(pattern) != len(s_list):
            return False

        # keep two dicts: one for pattern: s, and other for s: pattern
        # use it as a way to make both elements unique
        content = dict()
        content_s = dict()
        i = 0
        while i < len(pattern):
            left_pattern = pattern[i]
            left_s = s_list[i]
            if left_pattern in content:
                if content[left_pattern] != left_s:
                    # check that the character of the pattern is not already
                    # assigned to a different value of s
                    return False
            if left_s in content_s: 
                if content_s[left_s] != left_pattern:
                    # check that the value of s is not already
                    # assigned to a different character of the pattern
                    return False
            content[left_pattern] = left_s
            content_s[left_s] = left_pattern
            i += 1

        return True
