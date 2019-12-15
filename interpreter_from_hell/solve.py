flag = [0] * 38

flag[2] = 0x61 ^ 0x0c
flag[1] = 0x6c ^ 0x0b
flag[3] = 0x67 ^ 0x1e
flag[37] = 0x28 ^ 0x55
flag[0] = 0x66 ^ 0x11
flag[4] = 0x3a ^ 0x41

seed = flag[3] ^ (flag[2] << 8) ^ (flag[1] << 16) ^ (flag[0] << 24)

# get the seed to be used in random.cpp
# print("seed = ", seed)

from randomvalues import r

target = [160, 11, 119, 241, 178, 75, 110, 99, 239, 253, 170, 142, 217, 206, 80, 156]

# loop in reverse
for i in range(125 - 1, 0 - 1, -1):
    for j in range(16 - 1, 1 - 1, -1):
        target[j] ^= target[j-1]
    for j in range(16):
        target[j] ^= r[i][j]

for i in range(5, 21):
    flag[i] = target[i - 5]


flag[26] = flag[11]
"""
# commented because not needed, answers are at the bottom.

from z3 import BitVec, solve

# this also contains the parts we already defined above, which is going to be used later in gg
z = flag[:21] + [BitVec(f'flag{i}', 8) for i in range(21, 37)] + [flag[37]]

# conditions
c = []

c.append(z[26] == flag[26])
c.append(z[35] ^ z[36] ^ z[31] ^ z[34] == 1)
c.append(z[26] ^ z[36] ^ z[35] ^ z[21] == 81)
c.append(z[31] ^ z[22] ^ z[23] ^ z[27] == 85)
c.append(z[30] ^ z[25] ^ z[22] ^ z[34] == 6)
c.append(z[21] ^ z[29] ^ z[24] ^ z[26] == 7)
c.append(z[25] ^ z[23] ^ z[36] == 108)
c.append(z[36] ^ z[35] ^ z[25] == 51)
c.append(z[29] ^ z[32] ^ z[33] ^ z[21] == 80)
c.append(z[25] ^ z[26] ^ z[30] ^ z[34] == 6)
c.append(z[21] ^ z[24] ^ z[34] == 48)
c.append(z[29] ^ z[35] ^ z[30] ^ z[27] == 11)
c.append(z[34] ^ z[32] ^ z[23] ^ z[30] == 6)
c.append(z[33] ^ z[23] ^ z[26] ^ z[35] == 95)
c.append(z[32] ^ z[33] ^ z[30] == 98)
c.append(z[27] ^ z[28] ^ z[23] ^ z[30] == 2)

gg = 0

# this is from the beginning of the code, which checks a whole has or checksum.
for i in range(len(z)):
    gg ^= (z[i] << (i % 32))

c.append(gg == 0xb5973a46)

solve(*c)
"""
# put manually here after solving the z3 above, ranging from 21 to 36
solved = [49,
 97,
 102,
 98,
 50,
 97,
 49,
 99,
 53,
 54,
 99,
 53,
 97,
 99,
 57,
 56]


for i in range(21, 37):
    flag[i] = solved[i - 21]

print(bytes(flag))
