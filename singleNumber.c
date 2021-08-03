/*
Use: find the single not repeated number
Input: array and its length
Output: the singleNumber
 */
int singleNumber(int* nums, int numsSize)
{
    int i, singleNumber = NULL; 
    for (i = 0; i < numsSize; i++)
    {
        singleNumber ^= nums[i];
    }
    
    return singleNumber;
}
