/*
Use: check if a given array have duplicates
Input: array and its length
Output: if it has duplicates or not
 */
bool containsDuplicate(int* nums, int numsSize)
{
    int* countedNums = (int*)malloc(sizeof(int) * numsSize);
    int i, j, countedLen = 0;

    for (i = 0; i < numsSize; i ++)
    {
        for (j = 0; j < countedLen; j++)
        {
            if (countedNums[j] == nums[i])
            {
                return true;
            }
        }
        countedNums[i] = nums[i];
        countedLen++;
        
    }

    return false;
    
}

/*This is a more efficient solution*/

static int _cmp(const int *l, const int *r)  
{ 
    return (*l - *r); 
}

bool containsDuplicate(int* nums, int numsSize) {
    // sort 
    qsort(nums, numsSize, sizeof(int), _cmp);
    // check for repeat items
    for (int i = 1; i < numsSize; ++i) {
        if (nums[i] == nums[i - 1]) return true;
    }
    
    return false;
}
