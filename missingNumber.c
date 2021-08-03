/*
Use: find the missing number in array with range [0, n], n = array length
Input: array pointer, and the array size
Output: the missing number
 */
int missingNumber(int* nums, int n)
{
    int i, arraySum = 0, rangeSum = (n*(n + 1)) / 2;

    for (i = 0; i < n; i++)
    {
        arraySum += nums[i];
    }

    return rangeSum - arraySum;
}
