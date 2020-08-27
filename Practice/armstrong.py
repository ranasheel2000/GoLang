num = input("Enter a number: ")
l=len(num)
num=int(num)
sum = 0

temp = num
while temp > 0:
   digit = temp % 10
   sum += digit ** l
   temp //= 10

if num == sum:
   print(num,"is an Armstrong number")
else:
   print(num,"is not an Armstrong number")

