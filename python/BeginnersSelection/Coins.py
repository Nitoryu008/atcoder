a = int(input())
b = int(input())
c = int(input())
x = int(input())
count = 0

for ai in range(a + 1):
    for bi in range(b + 1):
        for ci in range(c + 1):
            if ai * 500 + bi * 100 + ci * 50 == x:
                count += 1

print(count)
