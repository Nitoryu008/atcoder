n = int(input())
a = list(map(int, input().split()))
count = 0

while True:
    isAllEven = True
    for ai in a:
        if ai % 2 != 0:
            isAllEven = False
            break
    if isAllEven:
        count += 1
        for i in range(n):
            a[i] /= 2
    else:
        break

print(count)
