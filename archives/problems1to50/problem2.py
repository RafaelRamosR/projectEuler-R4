#-------------------------------------------------------------------------------
# Name:        Even Fibonacci numbers
# Purpose:     By considering the terms in the Fibonacci sequence whose values
#              do not exceed four million, find the sum of the even-valued terms.
#
# Author:      Rafael R. Ramos R.
#
# Created:     19/05/2020
# Copyright:   (c) RRRR 2020
# Licence:     MIT
#-------------------------------------------------------------------------------

def fibonacci(limit):
    fibonacci, result, n = 0, 0, 0
    auro = (1 + 5**(0.5))/2

    while (fibonacci < limit):
        fibonacci = round((auro**n - (1 - auro)**n) / 5**(0.5))

        if fibonacci % 2 == 0:
            result = result + fibonacci

        n = n + 1
    return result

def problem2():
    print('Result:', fibonacci(4000000))