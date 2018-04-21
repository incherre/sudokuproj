# sudokuproj
A program written in golang for generating completed sudoku boards of a particular variant.

## Variant
### Example Board:
>            ,----2 4 5
>           /  ,--1 8 6
>           | | ,-3 7 9
>      ,----8 9 5 7 1 2
>     /  ,--6 4 2 8 9 3
>     | | ,-7 3 1 6 5 4
>     6 5 9 1 2 8 4 3 7
>     4 7 8 3 5 6 9 2 1
>     3 1 2 9 7 4 5 6 8   

### Rules
- Only the 6 3x3 squares on the lower right are used.
- Rules for 3X3 squares and complete rows/columns are the same as regular sudoku.
- Incomplete rows/columns are completed by following the curved line to a column/row of the size of the remaining elements.
Example: Checking the leftmost column, starting at the bottom, it goes 3 4 6 then curves into the row 4th from the top and 8 9 5 7 1 2.
