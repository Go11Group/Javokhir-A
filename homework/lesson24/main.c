#include<stdio.h>
#include<stdbool.h>
#include "puzzles.h"
#include<stdlib.h>
#include<unistd.h>

#define N 9

int puzzle[N][N] = {0};

void print_puzzle();
bool is_valid(int row, int col, int num);
void display_menu();
void settings();


int main(){
    easy_puzzle(puzzle);
    display_menu();
    
}

void play_game(){
    int row, col, num, user_choice;

    while (1)
    {
        system("clear");

        print_puzzle();

        puts("");
        printf("0.Back\nRow: ");
        scanf("%d", &row);
        if (row == 0){
            return;
        }
        printf("Column: ");
        scanf("%d", &col);
        printf("Num: ");
        scanf("%d", &num);

        row--;
        col--;

        if ((row >= 0 && row < 9) && (col >= 0 && col < 9)){
            if (is_valid(row, col, num)){
                puzzle[row][col] = num;
            }else{
                system("clear"); // Clear the screen before printing the error message
                print_puzzle(); // Print the puzzle again after clearing the screen
                printf("\nSame number found! Try again\n\n");
            }
        }else{
            system("clear"); // Clear the screen before printing the error message
            print_puzzle(); // Print the puzzle again after clearing the screen
            printf("\nInvalid index! Try again\n\n");
        }
        sleep(2);
    }

}

// for changing game settings
void settings(){
    system("clear");
    int user_choice;

    while (1){
        system("clear");
        printf("1.Set difficulty\n");
        printf("2.Change color\n");
        printf("3.Back\n");

        scanf("%d", &user_choice);

        switch (user_choice)
        {
        case 1:
            int mode;
            system("clear");
            printf("1.Easy Mode\n2.Medium Mode\n3.Hard Mode\n");
            scanf("%d", &mode);

            // setting puzzel board according to user's preference
            if (mode == 1){
                easy_puzzle(puzzle);
                printf("Mode changed to Easy\n");
            }else if(mode == 2){
                medium_puzzle(puzzle);
                printf("Mode changed to Medium\n");
            }else if(mode == 3){
                hard_puzzle(puzzle);
                printf("Mode changed to Hard\n");
            }else{
                printf("Invalid mode chosen!\n");
            }
            sleep(1);
            break;
        case 2:
            break;
        case 3: 
            return;
        default:
            printf("Invalid index!");
            break;
        }
    }

}

void display_menu(){
    int user_choice;
    while (1)
    {
        system("clear");
        printf("1.Start game\n");
        printf("2.Settings\n");
        printf("3.Quit\n");

        scanf("%d", &user_choice);

        switch (user_choice)
        {
        case 1:
            play_game();
            break;
        case 2:
            settings();
            break;
        case 3:
            exit(1);
        default:
            break;
        }
    }
    
}

void print_puzzle() {
    printf("   1 2 3 | 4 5 6 | 7 8 9  \n");
    printf("  -----------------------\n");
    for (int i = 0; i < N; i++) {
        printf("%d| ", i+1);
        for (int j = 0; j < N; j++) {
            if (puzzle[i][j] == 0) {
                printf(". ");
            } else {
                printf("%d ", puzzle[i][j]);
            }
            if ((j + 1) % 3 == 0) {
                printf("| ");
            }
        }
        printf("\n");
        if ((i + 1) % 3 == 0) {
            printf("  -----------------------\n");
        }
    }
}

bool is_valid(int row, int col, int num){
    for (int i = 0; i < 9; i++){
        if(puzzle[row][i] == num || puzzle[i][col] == num){
            return false;
        }
    }

    int startRow = row - row % 3;
    int startCol = col - col % 3;

    for(int i = 0; i < 3; i++){
        for (int j = 0; j < 3; j++){
            if (puzzle[i + startRow][j + startCol] == num){
                return false;
            }
        }
    }

    return true;
}