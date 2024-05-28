#include "puzzles.h"

void easy_puzzle(int puzzle[9][9]) {
    int puzzle1[9][9] = {
        {5, 0, 4, 6, 7, 8, 9, 1, 2},
        {6, 7, 0, 1, 0, 5, 0, 0, 8},
        {1, 9, 8, 0, 4, 2, 5, 6, 7},
        {0, 5, 9, 7, 0, 0, 4, 2, 3},
        {4, 2, 0, 0, 5, 3, 7, 0, 1},
        {7, 0, 3, 9, 2, 0, 8, 5, 6},
        {9, 6, 1, 0, 3, 7, 2, 0, 4},
        {2, 0, 7, 4, 1, 0, 6, 3, 0},
        {0, 4, 5, 0, 8, 6, 0, 7, 9}
    };
    
    for (int i = 0; i < 9; i++) {
        for (int j = 0; j < 9; j++) {
            puzzle[i][j] = puzzle1[i][j];
        }
    }
}

void medium_puzzle(int puzzle[9][9]) {
    int puzzle2[9][9] = {
        {0, 2, 3, 0, 5, 0, 7, 0, 9},
        {4, 5, 6, 7, 0, 9, 0, 2, 0},
        {0, 8, 0, 1, 2, 0, 4, 0, 6},
        {2, 0, 4, 0, 6, 7, 0, 9, 1},
        {5, 6, 7, 0, 9, 0, 2, 0, 4},
        {0, 0, 0, 0, 3, 4, 0, 6, 7},
        {3, 4, 5, 6, 7, 0, 9, 1, 2},
        {0, 7, 0, 0, 1, 2, 0, 0, 5},
        {9, 0, 2, 3, 4, 0, 6, 7, 0}
    };

    for (int i = 0; i < 9; i++) {
        for (int j = 0; j < 9; j++) {
            puzzle[i][j] = puzzle2[i][j];
        }
    }

}

void hard_puzzle(int puzzle[9][9]) {
    int puzzle3[9][9] = {
        {0, 8, 0, 0, 0, 1, 0, 0, 0},
        {2, 0, 9, 6, 0, 4, 0, 7, 1},
        {1, 0, 4, 0, 2, 0, 6, 0, 0},
        {0, 1, 0, 0, 0, 7, 0, 0, 4},
        {7, 0, 3, 0, 6, 0, 5, 0, 8},
        {0, 6, 5, 1, 0, 9, 0, 2, 0},
        {3, 5, 0, 8, 9, 0, 1, 0, 2},
        {0, 0, 1, 0, 0, 3, 0, 9, 5},
        {9, 0, 8, 0, 0, 0, 3, 0, 0}
    };

    
    for (int i = 0; i < 9; i++) {
        for (int j = 0; j < 9; j++) {
            puzzle[i][j] = puzzle3[i][j];
        }
    }
}