
class GFG{

/* Function to calculate (base^exponent)%modulus */
static long  modular_pow(long  base, int exponent,
                          long  modulus)
{
    /* initialize result */
    long  result = 1;

    while (exponent > 0)
    {
        /* if y is odd, multiply base with result */
        if (exponent % 2 == 1)
            result = (result * base) % modulus;

        /* exponent = exponent/2 */
        exponent = exponent >> 1;

        /* base = base * base */
        base = (base * base) % modulus;
    }
    return result;
}


// Recursive function to return gcd of a and b
static long __gcd(long a, long b)
{
 return b == 0? a:__gcd(b, a % b);
}

/* driver function */
public static void main(String[] args)
    {
    }
}
