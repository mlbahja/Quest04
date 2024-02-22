#include <stdio.h>
#include <unistd.h>

#define SIZE 10

void ft_putchar(char c)
{
    write(1, &c, 1);
}

void print_position(int *board)
{
    int i;

    i = 0;
    while (i < SIZE)
        ft_putchar(board[i++] + 48);
    ft_putchar('\n');
}

int is_valid(int *board, int row, int col)
{
    int i;

    i = 0;
    while (i < row)
    {
        if (board[i] == col || board[i] == col - row + i || board[i] == col + row - i)
            return (0);
        i++;
    }
    return (1);
}

int put_queen(int *board, int row, int *pos)
{
    int col;

    if (row == SIZE)
    {
        (*pos)++;
        print_position(board);
    }
    else
    {
        col = 0;
        while (col < SIZE)
        {
            if (is_valid(board, row, col))
            {
                board[row] = col;
                put_queen(board, row + 1, pos);
            }
            col++;
        }
    }
    return (*pos);
}

int ft_ten_queens_puzzle(void)
{
    int board[SIZE] = {0}; // Initialize the board array with zeros
    int row = 0;
    int pos = 0;

    return put_queen(board, row, &pos);
}

int main()
{
    printf("Number of solutions: %d\n", ft_ten_queens_puzzle());
    return 0;
}
