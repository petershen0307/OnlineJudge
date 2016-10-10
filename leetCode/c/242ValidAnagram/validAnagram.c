#include <string.h>
#include <stdbool.h>

#define array_size 256
void check_hash(char const *s, char *hash, bool increased)
{
    while(0 != *s)
    {
        if (array_size <= *s || 0 > *s)
        {
            return; //error
        }
        if (increased)
        {
            ++hash[*s];
        }
        else
        {
            --hash[*s];
        }
        ++s;
    }
}

bool isAnagram(char* s, char* t) {
    char hash[array_size];
    memset(hash, 0, array_size);
    check_hash(s, hash, true);
    check_hash(t, hash, false);
    for (int i = 0; i < array_size; ++i)
    {
        if (0 != hash[i])
        {
            return false;
        }
    }
    return true;
}
