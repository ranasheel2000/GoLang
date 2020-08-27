def compute_hcfGcd(x, y):
    if x > y:
        smaller = y
    else:
        smaller = x
    hcf=0
    for i in range(smaller,1,-1):
        if((x % i == 0) and (y % i == 0)):
            hcf = i
            break#return(hcf)
    return(hcf)

num1 = 35 
num2 = 45
print("hcf/gcd is", compute_hcfGcd(num1, num2))
