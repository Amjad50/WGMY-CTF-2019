# interpreter_from_hell <!-- omit in toc -->

```
We asked our junior to make one programming language for us. He went too far...
```

```
$ file interpreter code.txt 
interpreter: ELF 64-bit LSB executable, x86-64, version 1 (GNU/Linux), statically linked, for GNU/Linux 3.2.0, not stripped
code.txt:    ASCII text
```

- [Introduction](#introduction)
- [Information collection](#information-collection)
  - [Interpreter analysis](#interpreter-analysis)
    - [Interpreter::main](#interpretermain)
    - [Interpreter::run_program](#interpreterrunprogram)
- [Solution](#solution)
  - [solution process](#solution-process)
    - [step 5](#step-5)
    - [step 7](#step-7)


## Introduction
This challenge introduce an interpreter with a code that works with the interpreter.

**Interpreters:** are programs that read piece of code, analyse, and run during run-time. Some of the interpreted languages: (python, JS, `and this (challenge)`)

In this challenge, the `interpreter` is written in `C++`. We are going to use decompilation tools like `Ghidra` to help us solve it.

Note: This challenge, took a lot of time from me, as I did a lot of manual work :(.

## Information collection
First, we look at [code.txt], and its in a very bad shape. WHAT is this language??

Somehow the `interpreter` binary is able to run that piece of code.

### Interpreter analysis
The good thing about this challenge is that `interpreter` is not stripped, which
makes our analysis easier, specially because its `C++`.

Looking at `main`, there is two functions in this program that are important:

#### Interpreter::main
Starting point of the program.

- Checks if the user passed a file argument
- reads each line of the file and store it into a `global` `vector` object `program_data`.
- initializes two `maps` objects `map<string, string>` and `map<string, vector<int>>`
- calls `run_program` and passes the maps that were created. And passes the number of the first line and the last line to execute from the code (the reason will come later).

#### Interpreter::run_program
The fun is here.

This function has a gigantic loop, to loop from the starting line and ending line passed to the function.

In the loop:
- split by `' '` and select the first part of each line, which is the command.
- This command goes to a very big `if else` tree (1400 lines).
- In the `if else` it checks the command with every command defined in this
`language`. 
- If the command matches it runs a piece of code to do the task of that command, 
which is how interpreters work in simple sense.

Here is the manual work, what I did is go through [code.txt] line by line, get the
first part (command) and looked what it does in `Ghidra` and record the function of
tha command.

The modefied **(readable)** code can be found in [code_edited.txt]
Made the names of the commands self explanatory.

As stated above, that this interpreter has two `maps`. These are used to store
variables in this programming language. it can only store `string` or
`vector<int>`, in the program `vector<int>` is used as a `char array`.

Even commands like `assignNumber` is using the string map to store the number
as a string.

> should I group the commands into a file to tell what they do??

## Solution
After decoding the code

Here is what the code do:
1) read a line from the user
2) the line length must be 38 chars
3) the hash/checksum of the line must match `0xb5973a46`
4) using XOR to check that the flag has `wgmy{*}`
5) using encryption to check the middle part of the flag [5:21]
6) check that line[11] is equal to line[26]
7) series of XORs to determine that line[21:37] are OK
8) If all the above is OK, then this is the correct flag.

we have all the specs for the flag, we can put them together

### solution process
I will leave **(3)** until the end, **(1), (2) and (4)** are easy.

#### step 5
As for **(5)** its in `code_edited.txt` lines 72 - 150.

What it does is
1) copies the input from 5 - 21 into a new vector
2) generate a seed value from `"wgmy"` by using shift and XOR.
3) using the seed to initialize random
4) do a series of XORing with the random values and XORing among the flag
characters.

One command to pay attention to is `randomPushRandomSplit4 random_vector`,
which generate one int32 random number, split it into 4 bytes and store the
bytes in the argument vector variable (`random_vector`).

Then `vectorAtIndexCopy random_vector 3 random_num` is executed to get only
the last byte of the random number, thus we only need one byte of each random
generation operation.

The series of operations (125 times):
1) XOR each character [5 - 21] with a new random number
2) XOR each character with the one before it

To generate the random values I used `C++` code [random.cpp](random.cpp)
, and stored the result into [randomvalues.py](randomvalues.py) to be
used to solve for that part of the flag.

```python
# from solve.py
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
```

#### step 7
step **(6)** is straight forward, but important for this step **(7)**

step **(7)** is taking part from code_edited.txt lines 161 - 316

In this step, the program do some XORs between characters [21 - 36] and
compare the results with predefined values.

I used `z3-solver` to solve this part.

I collected all cases **manually** and put them here

```python
from z3 import BitVec, solve

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
```

During the CTF, I tried this but didn't work and then I realized that there is
step **(3)** which computes the checksum for the whole flag, so I added it
also.

```python
gg = 0

# checksum
for i in range(len(z)):
    gg ^= (z[i] << (i % 32))

c.append(gg == 0xb5973a46)
```

When running the above code we get the following values for the result. 

```python
[flag25 = 50,
 flag30 = 54,
 flag31 = 99,
 flag24 = 98,
 flag33 = 97,
 flag23 = 102,
 flag35 = 57,
 flag34 = 99,
 flag32 = 53,
 flag28 = 99,
 flag21 = 49,
 flag22 = 97,
 flag36 = 56,
 flag27 = 49,
 flag29 = 53,
 flag26 = 97]
```

When compiling everything we get the flag:
```
wgmy{10e3b9a7cb5a88391afb2a1c56c5ac98}
```


[code.txt]: challenge&#32;files/code.txt
[code_edited.txt]: code_edited.txt