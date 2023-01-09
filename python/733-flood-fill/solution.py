"""
The keys for the Flood Fill are:
    - Get valid neighbors: [up, down , left, right] -> within the matrix dimensions
    - Color of the source pixel
    - If source pixel color and color are the same: return image

1. Iterate over each of the neighbors
2. If the color matches the source pixel, update color
2. Call recursively for the neighbors of the neighbors, pass updated image matrix
"""
from typing import List

class Solution:
    moves = [
        [-1, 0],
        [1, 0],
        [0, 1],
        [0, -1],
    ]

    def fill(self, image, sr, sc, color, def_color, m, n):
        for move in self.moves:
            r = sr+move[0]
            c = sc+move[1]
            if 0 <= r < m and 0 <= c < n:
                if image[r][c] == def_color:
                    image[r][c] = color
                    self.fill(image, r, c, color, def_color, m, n)

        return image


    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        m = len(image)
        n = len(image[0])
        def_color = image[sr][sc]
        if color == def_color:
            return image
        image[sr][sc] = color
        result = self.fill(image, sr, sc, color, def_color, m, n)

        return result

if __name__ == "__main__":
    sol = Solution()
    image1 = [
            [1,1,1],
            [1,1,0],
            [1,0,1],
            ]
    sr, sc = 1, 1
    color = 2
    result1 = sol.floodFill(image1, sr, sc, color)
    print("Result 1: ", result1)
    image2 = [
            [0,0,0],
            [0,0,0],
            ]
    sr, sc = 0, 0
    color = 0
    result2 = sol.floodFill(image2, sr, sc, color)
    print("Result 2: ", result2)
