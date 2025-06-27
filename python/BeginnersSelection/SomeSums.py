n, a, b = map(int, input().split())
total = 0

for i in range(1, n + 1):
    sum = i % 10 + i % 100 // 10 + i % 1000 // 100 + i % 10000 // 1000 + i // 10000
    if a <= sum <= b:
        total += i

print(total)
