def swap(numbers):
    for i in range(len(numbers)-1, 0, -1):
        for j in range(i):
            if numbers[j]<numbers[j+1]:
                numbers[j],numbers[j+1] = numbers[j+1],numbers[j]
                
numbers = [21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12]
print(numbers)
swap(numbers)
print(numbers)

