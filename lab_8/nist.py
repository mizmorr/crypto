import string

import numpy

path ="sample/mt_10000.txt"
with open(path) as file:
    text = file.read()


import matplotlib.pyplot as plt


# m = 25

# freqs = []
# for k in range (0,len(text)-m,m):
#     for letter in {'1'}:
#         letter_freq = text[k:k+m].count(letter)
#         freqs.append(letter_freq/m)

freqs = []
for digit in string.digits:
    digit_freq = text.count(digit)
    print(digit,digit_freq)
    freqs.append(digit_freq)

n, bins, patches = plt.hist(freqs)
plt.show()
# print("mean - ", numpy.mean(freqs))
