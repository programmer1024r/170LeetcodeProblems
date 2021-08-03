/*
Use: remove an element form a given array
Input: array his length, and the index to remove
Output: None
 */
void removeElement(int* array, int n, int index)
{
    int i;
    for (i = index; i < n - 1; i++)
    {
        // shift
        array[i] = array[i + 1];
    }
    
}

/*
Use: find the missing numbers in array with range [1, n], n = array length
Input: array pointer, and the array size
Output: the missing numbers
 */
int* findDisappearedNumbers(int* nums, int n, int* solution)
{
    int i;
    solution = (int*)malloc(sizeof(int) * n);

    if (n == 0) return 0;

    for (i = 0; i < n; i++)
    {
        solution[i] = i + 1;
    }
    
    // go throw the array and remove all 
    // the numbers that appear in it from the solution
    for (i = 0; i < n; i++)
    {
        // if the element we want to remove is not its duplicate
        if (solution[nums[i] - 1] == nums[i])
        {
            // Remove
            removeElement(solution, n, nums[i - 1]);
        }
    }

    return solution;
}

